
# Table: k8s_networking_network_policy_egress_tos
NetworkPolicyPeer describes a peer to allow traffic to/from
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|network_policy_egress_cq_id|uuid|Unique CloudQuery ID of k8s_networking_network_policy_egresses table (FK)|
|pod_selector_match_labels|jsonb||
|namespace_selector_match_labels|jsonb||
|ip_block_c_id_r|text|CIDR is a string representing the IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64"|
|ip_block_except|text[]|Except is a slice of CIDRs that should not be included within an IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64" Except values will be rejected if they are outside the CIDR range +optional|
