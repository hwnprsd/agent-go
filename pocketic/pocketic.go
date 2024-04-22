package pocketic

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/aviate-labs/agent-go/principal"
	"time"
)

var DefaultSubnetConfig = ExtendedSubnetConfigSet{
	NNS: &SubnetSpec{
		StateConfig:       NewSubnetStateConfig{},
		InstructionConfig: ProductionSubnetInstructionConfig{},
		DtsFlag:           false,
	},
}

// BenchmarkingSubnetInstructionConfig uses very high instruction limits useful for asymptotic canister benchmarking.
type BenchmarkingSubnetInstructionConfig struct{}

func (c BenchmarkingSubnetInstructionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal("Benchmarking")
}

func (c BenchmarkingSubnetInstructionConfig) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}
	if s != "Benchmarking" {
		return fmt.Errorf("invalid instruction config: %s", s)
	}
	return nil
}

func (BenchmarkingSubnetInstructionConfig) instructionConfig() {}

type CanisterSettings struct {
	Controllers       *[]principal.Principal `ic:"controllers,omitempty" json:"controllers,omitempty"`
	ComputeAllocation *idl.Nat               `ic:"compute_allocation,omitempty" json:"compute_allocation,omitempty"`
	MemoryAllocation  *idl.Nat               `ic:"memory_allocation,omitempty" json:"memory_allocation,omitempty"`
	FreezingThreshold *idl.Nat               `ic:"freezing_threshold,omitempty" json:"freezing_threshold,omitempty"`
}

type CreateCanisterArgs struct {
	Settings    *CanisterSettings    `ic:"settings,omitempty" json:"settings,omitempty"`
	SpecifiedID *principal.Principal `ic:"specified_id" json:"specified_id,omitempty"`
}

type DtsFlag bool

func (f DtsFlag) MarshalJSON() ([]byte, error) {
	if f {
		return json.Marshal("Enabled")
	}
	return json.Marshal("Disabled")
}

func (f *DtsFlag) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}
	if s != "Enabled" && s != "Disabled" {
		return fmt.Errorf("invalid DTS flag: %s", s)
	}
	*f = s == "Enabled"
	return nil
}

type EffectiveCanisterID struct {
	CanisterId string `json:"CanisterId"`
}

type EffectiveSubnetID struct {
	SubnetID string `json:"SubnetId"`
}

// FromPathSubnetStateConfig load existing subnet state from the given path. The path must be on a filesystem
// accessible to the server process.
type FromPathSubnetStateConfig struct {
	Path     string
	SubnetID RawSubnetID
}

func (c FromPathSubnetStateConfig) UnmarshalJSON(bytes []byte) error {
	var v []json.RawMessage
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if len(v) != 2 {
		return fmt.Errorf("invalid state config: %v", v)
	}
	if err := json.Unmarshal(v[0], &c.Path); err != nil {
		return err
	}
	return json.Unmarshal(v[1], &c.SubnetID)
}

func (c FromPathSubnetStateConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal([]any{c.Path, c.SubnetID})
}

func (FromPathSubnetStateConfig) stateConfig() {}

type NNSConfig struct {
	StateDirPath string
	SubnetID     principal.Principal
}

// NewSubnetStateConfig creates new subnet with empty state.
type NewSubnetStateConfig struct{}

func (c NewSubnetStateConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal("New")
}

func (c NewSubnetStateConfig) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}
	if s != "New" {
		return fmt.Errorf("invalid state config: %s", s)
	}
	return nil
}

func (NewSubnetStateConfig) stateConfig() {}

type PocketIC struct {
	server     *server
	instanceID int
	topology   map[string]Topology
	sender     principal.Principal
}

// New creates a new PocketIC instance with the given subnet configuration.
func New(subnetConfig ExtendedSubnetConfigSet) (*PocketIC, error) {
	s, err := newServer()
	if err != nil {
		return nil, err
	}
	resp, err := s.NewInstance(subnetConfig)
	if err != nil {
		return nil, err
	}
	return &PocketIC{
		server:     s,
		instanceID: resp.InstanceID,
		topology:   resp.Topology,
		sender:     principal.AnonymousID,
	}, nil
}

func (pic PocketIC) AddCycles(canisterID principal.Principal, amount int) (int, error) {
	var body struct {
		Cycles int `json:"cycles"`
	}
	if err := pic.server.InstancePost(pic.instanceID, "update/add_cycles", map[string]any{
		"canister_id": base64.StdEncoding.EncodeToString(canisterID.Raw),
		"amount":      amount,
	}, &body); err != nil {
		return 0, err
	}
	return body.Cycles, nil
}

// AdvanceTime advances the time of the PocketIC instance by the given nanoseconds.
func (pic PocketIC) AdvanceTime(nanoSeconds int) error {
	t, err := pic.GetTime()
	if err != nil {
		return err
	}
	return pic.server.InstancePost(pic.instanceID, "update/set_time", map[string]any{
		"nanos_since_epoch": t.Nanosecond() + nanoSeconds,
	}, nil)
}

// CanisterExits returns true if the given canister exists in the PocketIC instance.
func (pic PocketIC) CanisterExits(canisterID principal.Principal) bool {
	_, err := pic.GetSubnet(canisterID)
	return err == nil
}

func (pic PocketIC) CreateAndInstallCanister(wasmModule []byte, arg []byte, subnetPID *principal.Principal) (*principal.Principal, error) {
	canisterID, err := pic.CreateCanister(CreateCanisterArgs{}, subnetPID)
	if err != nil {
		return nil, err
	}
	if _, err := pic.AddCycles(*canisterID, 2_000_000_000_000); err != nil {
		return nil, err
	}
	if err := pic.InstallCode(*canisterID, wasmModule, arg); err != nil {
		return nil, err
	}
	return canisterID, nil
}

func (pic PocketIC) CreateCanister(args CreateCanisterArgs, subnetPID *principal.Principal) (*principal.Principal, error) {
	var ecID any
	if subnetPID != nil {
		ecID = EffectiveSubnetID{
			SubnetID: base64.StdEncoding.EncodeToString(subnetPID.Raw),
		}
	}

	payload, err := idl.Marshal([]any{args})
	if err != nil {
		return nil, err
	}

	var resp struct {
		CanisterID principal.Principal `ic:"canister_id"`
	}
	if err := pic.UpdateCallWithEffectiveCanisterID(
		nil,
		ecID,
		"provisional_create_canister_with_cycles",
		payload,
		[]any{&resp},
	); err != nil {
		return nil, err
	}
	return &resp.CanisterID, nil
}

func (pic PocketIC) GetCycleBalance(canisterID principal.Principal) (int, error) {
	var body struct {
		Cycles int `json:"cycles"`
	}
	if err := pic.server.InstancePost(pic.instanceID, "read/get_cycles", map[string]any{
		"canister_id": base64.StdEncoding.EncodeToString(canisterID.Raw),
	}, &body); err != nil {
		return 0, err
	}
	return body.Cycles, nil
}

// GetRootKey returns the root key of the PocketIC instance.
func (pic PocketIC) GetRootKey() ([]byte, error) {
	var nnsPID principal.Principal
	for k, v := range pic.topology {
		if v.SubnetKind == NNSSubnet {
			pid, err := principal.Decode(k)
			if err != nil {
				return nil, err
			}
			nnsPID = pid
			break
		}

	}
	if len(nnsPID.Raw) == 0 {
		return nil, fmt.Errorf("NNS subnet not found")
	}
	var body []byte
	if err := pic.server.InstancePost(pic.instanceID, "read/pub_key", map[string]any{
		"subnet_id": base64.StdEncoding.EncodeToString(nnsPID.Raw),
	}, &body); err != nil {
		return nil, err
	}
	return body, nil
}

// GetSubnet returns the subnet of the given canister.
func (pic PocketIC) GetSubnet(canisterID principal.Principal) (*principal.Principal, error) {
	var body struct {
		SubnetID string `json:"subnet_id"`
	}
	if err := pic.server.InstancePost(pic.instanceID, "read/get_subnet", map[string]any{
		"canister_id": base64.StdEncoding.EncodeToString(canisterID.Raw),
	}, &body); err != nil {
		return nil, err
	}
	subnetPID, err := base64.StdEncoding.DecodeString(body.SubnetID)
	if err != nil {
		return nil, err
	}
	return &principal.Principal{
		Raw: subnetPID,
	}, nil
}

// GetTime returns the current time of the PocketIC instance.
func (pic PocketIC) GetTime() (*time.Time, error) {
	var m struct {
		NanosSinceEpoch int64 `json:"nanos_since_epoch"`
	}
	if err := pic.server.InstanceGet(pic.instanceID, "read/get_time", &m); err != nil {
		return nil, err
	}
	t := time.Unix(0, m.NanosSinceEpoch)
	return &t, nil
}

func (pic PocketIC) InstallCode(canisterID principal.Principal, wasmModule []byte, arg []byte) error {
	payload, err := idl.Marshal([]any{installCodeArgs{
		WasmModule: wasmModule,
		CanisterID: canisterID,
		Arg:        arg,
		Mode: installMode{
			Install: &idl.Null{},
		},
	}})
	if err != nil {
		return err
	}
	return pic.UpdateCallWithEffectiveCanisterID(
		nil,
		EffectiveCanisterID{
			CanisterId: base64.StdEncoding.EncodeToString(canisterID.Raw),
		},
		"install_code",
		payload,
		nil,
	)
}

func (pic PocketIC) QueryCall(canisterID principal.Principal, method string, payload []any, body []any) error {
	rawPayload, err := idl.Marshal(payload)
	if err != nil {
		return err
	}
	return pic.canisterCall("read/query", &canisterID, nil, method, rawPayload, body)
}

// SetSender sets the sender principal for the PocketIC instance.
func (pic *PocketIC) SetSender(sender principal.Principal) {
	pic.sender = sender
}

// SetTime sets the time of the PocketIC instance to the given nanoseconds since epoch.
func (pic PocketIC) SetTime(nanosSinceEpoch int) error {
	return pic.server.InstancePost(pic.instanceID, "update/set_time", map[string]any{
		"nanos_since_epoch": nanosSinceEpoch,
	}, nil)
}

// Tick advances the PocketIC instance by one block.
func (pic PocketIC) Tick() error {
	return pic.server.InstancePost(pic.instanceID, "update/tick", nil, nil)
}

func (pic PocketIC) UpdateCall(canisterID principal.Principal, method string, payload []any, body []any) error {
	rawPayload, err := idl.Marshal(payload)
	if err != nil {
		return err
	}
	return pic.UpdateCallWithEffectiveCanisterID(&canisterID, nil, method, rawPayload, body)
}

func (pic PocketIC) UpdateCallWithEffectiveCanisterID(canisterID *principal.Principal, ecID any, method string, payload []byte, body []any) error {
	return pic.canisterCall("update/execute_ingress_message", canisterID, ecID, method, payload, body)
}

func (pic PocketIC) canisterCall(endpoint string, canisterID *principal.Principal, ecID any, method string, payload []byte, body []any) error {
	if ecID == nil {
		ecID = "None"
	}
	var cID principal.Principal
	if canisterID != nil {
		cID = *canisterID
	}
	var reply reply
	if err := pic.server.InstancePost(pic.instanceID, endpoint, map[string]any{
		"sender":              base64.StdEncoding.EncodeToString(pic.sender.Raw),
		"effective_principal": ecID,
		"canister_id":         base64.StdEncoding.EncodeToString(cID.Raw),
		"method":              method,
		"payload":             base64.StdEncoding.EncodeToString(payload),
	}, &reply); err != nil {
		return err
	}
	if reply.Ok == nil {
		return reply.Err
	}
	if reply.Ok.Reply == nil {
		return *reply.Ok.Reject
	}
	rawBody, err := base64.StdEncoding.DecodeString(*reply.Ok.Reply)
	if err != nil {
		return err
	}
	return idl.Unmarshal(rawBody, body)
}

// ProductionSubnetInstructionConfig uses default instruction limits as in production.
type ProductionSubnetInstructionConfig struct{}

func (c ProductionSubnetInstructionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal("Production")
}

func (c ProductionSubnetInstructionConfig) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}
	if s != "Production" {
		return fmt.Errorf("invalid instruction config: %s", s)
	}
	return nil
}

func (ProductionSubnetInstructionConfig) instructionConfig() {}

type RawSubnetID struct {
	SubnetID string `json:"subnet_id"`
}

func (r RawSubnetID) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"subnet-id": base64.StdEncoding.EncodeToString([]byte(r.SubnetID)),
	})
}

func (r *RawSubnetID) UnmarshalJSON(bytes []byte) error {
	var rawSubnetID struct {
		SubnetID string `json:"subnet_id-id"`
	}
	if err := json.Unmarshal(bytes, &rawSubnetID); err != nil {
		return err
	}
	subnetID, err := base64.StdEncoding.DecodeString(rawSubnetID.SubnetID)
	if err != nil {
		return err
	}
	r.SubnetID = string(subnetID)
	return nil
}

type ExtendedSubnetConfigSet struct {
	Application []SubnetSpec `json:"application"`
	Bitcoin     *SubnetSpec  `json:"bitcoin,omitempty"`
	Fiduciary   *SubnetSpec  `json:"fiduciary,omitempty"`
	II          *SubnetSpec  `json:"ii,omitempty"`
	NNS         *SubnetSpec  `json:"nns,omitempty"`
	SNS         *SubnetSpec  `json:"sns,omitempty"`
	System      []SubnetSpec `json:"system"`
}

type RejectError string

func (e RejectError) Error() string {
	return string(e)
}

type ReplyError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func (e ReplyError) Error() string {
	return fmt.Sprintf("code: %d, description: %s", e.Code, e.Description)
}

type SubnetInstructionConfig interface {
	instructionConfig()
}

type SubnetKind string

var (
	ApplicationSubnet SubnetKind = "Application"
	BitcoinSubnet     SubnetKind = "Bitcoin"
	FiduciarySubnet   SubnetKind = "Fiduciary"
	IISubnet          SubnetKind = "II"
	NNSSubnet         SubnetKind = "NNS"
	SNSSubnet         SubnetKind = "SNS"
	SystemSubnet      SubnetKind = "System"
)

// SubnetSpec specifies various configurations for a subnet.
type SubnetSpec struct {
	StateConfig       SubnetStateConfig       `json:"state_config"`
	InstructionConfig SubnetInstructionConfig `json:"instruction_config"`
	DtsFlag           DtsFlag                 `json:"dts_flag"`
}

type SubnetStateConfig interface {
	stateConfig()
}

type installCodeArgs struct {
	WasmModule []byte              `ic:"wasm_module"`
	CanisterID principal.Principal `ic:"canister_id"`
	Arg        []byte              `ic:"arg"`
	Mode       installMode         `ic:"mode"`
}

type installMode struct {
	Install   *idl.Null `ic:"install,variant"`
	Reinstall *idl.Null `ic:"reinstall,variant"`
	Upgrade   *idl.Null `ic:"upgrade,variant"`
}

type reply struct {
	Ok *struct {
		Reply  *string      `json:"Reply,omitempty"`
		Reject *RejectError `json:"Reject,omitempty"`
	} `json:"Ok,omitempty"`
	Err *ReplyError `json:"Err,omitempty"`
}
