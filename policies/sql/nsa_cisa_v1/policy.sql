\set framework 'nsa_cisa_v1'
\set execution_time `date '+%Y-%m-%d %H:%M:%S'`
\i sql/create_k8s_policy_results.sql
\i sql/nsa_cisa_v1/pod_security.sql
\i sql/nsa_cisa_v1/network_hardening.sql
