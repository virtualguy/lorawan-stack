// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var ListFrequencyPlansRequestFieldPathsNested = []string{
	"base_frequency",
}

var ListFrequencyPlansRequestFieldPathsTopLevel = []string{
	"base_frequency",
}
var FrequencyPlanDescriptionFieldPathsNested = []string{
	"base_frequency",
	"base_id",
	"id",
	"name",
}

var FrequencyPlanDescriptionFieldPathsTopLevel = []string{
	"base_frequency",
	"base_id",
	"id",
	"name",
}
var ListFrequencyPlansResponseFieldPathsNested = []string{
	"frequency_plans",
}

var ListFrequencyPlansResponseFieldPathsTopLevel = []string{
	"frequency_plans",
}
var GetPhyVersionsRequestFieldPathsNested = []string{
	"band_id",
}

var GetPhyVersionsRequestFieldPathsTopLevel = []string{
	"band_id",
}
var GetPhyVersionsResponseFieldPathsNested = []string{
	"version_info",
}

var GetPhyVersionsResponseFieldPathsTopLevel = []string{
	"version_info",
}
var ListBandsRequestFieldPathsNested = []string{
	"band_id",
	"phy_version",
}

var ListBandsRequestFieldPathsTopLevel = []string{
	"band_id",
	"phy_version",
}
var BandDescriptionFieldPathsNested = []string{
	"adr_ack_limit",
	"beacon",
	"beacon.coding_rate",
	"beacon.data_rate_index",
	"beacon.frequencies",
	"boot_dwell_time",
	"boot_dwell_time.downlinks",
	"boot_dwell_time.uplinks",
	"cf_list_type",
	"data_rates",
	"default_max_eirp",
	"default_rx2_parameters",
	"default_rx2_parameters.data_rate_index",
	"default_rx2_parameters.frequency",
	"downlink_channels",
	"freq_multiplier",
	"id",
	"implements_cf_list",
	"join_accept_delay_1",
	"join_accept_delay_2",
	"max_adr_data_rate_index",
	"max_downlink_channels",
	"max_fcnt_gap",
	"max_retransmit_timeout",
	"max_uplink_channels",
	"min_retransmit_timeout",
	"ping_slot_frequencies",
	"receive_delay_1",
	"receive_delay_2",
	"relay",
	"relay.wor_channels",
	"sub_bands",
	"supports_dynamic_adr",
	"tx_offset",
	"tx_param_setup_req_support",
	"uplink_channels",
}

var BandDescriptionFieldPathsTopLevel = []string{
	"adr_ack_limit",
	"beacon",
	"boot_dwell_time",
	"cf_list_type",
	"data_rates",
	"default_max_eirp",
	"default_rx2_parameters",
	"downlink_channels",
	"freq_multiplier",
	"id",
	"implements_cf_list",
	"join_accept_delay_1",
	"join_accept_delay_2",
	"max_adr_data_rate_index",
	"max_downlink_channels",
	"max_fcnt_gap",
	"max_retransmit_timeout",
	"max_uplink_channels",
	"min_retransmit_timeout",
	"ping_slot_frequencies",
	"receive_delay_1",
	"receive_delay_2",
	"relay",
	"sub_bands",
	"supports_dynamic_adr",
	"tx_offset",
	"tx_param_setup_req_support",
	"uplink_channels",
}
var ListBandsResponseFieldPathsNested = []string{
	"descriptions",
}

var ListBandsResponseFieldPathsTopLevel = []string{
	"descriptions",
}
var GetPhyVersionsResponse_VersionInfoFieldPathsNested = []string{
	"band_id",
	"phy_versions",
}

var GetPhyVersionsResponse_VersionInfoFieldPathsTopLevel = []string{
	"band_id",
	"phy_versions",
}
var BandDescription_BeaconFieldPathsNested = []string{
	"coding_rate",
	"data_rate_index",
	"frequencies",
}

var BandDescription_BeaconFieldPathsTopLevel = []string{
	"coding_rate",
	"data_rate_index",
	"frequencies",
}
var BandDescription_ChannelFieldPathsNested = []string{
	"frequency",
	"max_data_rate",
	"min_data_rate",
}

var BandDescription_ChannelFieldPathsTopLevel = []string{
	"frequency",
	"max_data_rate",
	"min_data_rate",
}
var BandDescription_SubBandParametersFieldPathsNested = []string{
	"duty_cycle",
	"max_eirp",
	"max_frequency",
	"min_frequency",
}

var BandDescription_SubBandParametersFieldPathsTopLevel = []string{
	"duty_cycle",
	"max_eirp",
	"max_frequency",
	"min_frequency",
}
var BandDescription_BandDataRateFieldPathsNested = []string{
	"rate",
	"rate.modulation",
	"rate.modulation.fsk",
	"rate.modulation.fsk.bit_rate",
	"rate.modulation.lora",
	"rate.modulation.lora.bandwidth",
	"rate.modulation.lora.coding_rate",
	"rate.modulation.lora.spreading_factor",
	"rate.modulation.lrfhss",
	"rate.modulation.lrfhss.coding_rate",
	"rate.modulation.lrfhss.modulation_type",
	"rate.modulation.lrfhss.operating_channel_width",
}

var BandDescription_BandDataRateFieldPathsTopLevel = []string{
	"rate",
}
var BandDescription_Rx2ParametersFieldPathsNested = []string{
	"data_rate_index",
	"frequency",
}

var BandDescription_Rx2ParametersFieldPathsTopLevel = []string{
	"data_rate_index",
	"frequency",
}
var BandDescription_DwellTimeFieldPathsNested = []string{
	"downlinks",
	"uplinks",
}

var BandDescription_DwellTimeFieldPathsTopLevel = []string{
	"downlinks",
	"uplinks",
}
var BandDescription_RelayParametersFieldPathsNested = []string{
	"wor_channels",
}

var BandDescription_RelayParametersFieldPathsTopLevel = []string{
	"wor_channels",
}
var BandDescription_RelayParameters_RelayWORChannelFieldPathsNested = []string{
	"ack_frequency",
	"data_rate_index",
	"frequency",
}

var BandDescription_RelayParameters_RelayWORChannelFieldPathsTopLevel = []string{
	"ack_frequency",
	"data_rate_index",
	"frequency",
}
var ListBandsResponse_VersionedBandDescriptionFieldPathsNested = []string{
	"band",
}

var ListBandsResponse_VersionedBandDescriptionFieldPathsTopLevel = []string{
	"band",
}
