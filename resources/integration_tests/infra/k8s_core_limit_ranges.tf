resource "kubernetes_limit_range" "example" {
  metadata {
    name = "limit-range${var.test_prefix}${var.test_suffix}"
  }

  spec {
    limit {
      type = "Pod"
      max = {
        cpu    = "200m"
        memory = "1024Mi"
      }
    }
  }
}
