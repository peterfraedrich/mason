cloud {
  aws {
    Ec2AmiID            = ""
    Ec2AvailabilityZone = ""
    Ec2InstanceID       = ""
    Ec2InstanceType     = ""
    Ec2Profile          = ""
    Ec2PublicIP4        = ""
    Ec2IAMID            = ""
    Ec2IAMARN           = ""
  }
}

cpu {
  cache_size = 512
  cores      = 16
  mhz        = 3593.298
  model      = "AMD Ryzen 7 1800X Eight-Core Processor"
  vendor     = "AuthenticAMD"
}

host {
  architecture   = "x86_64"
  boot_time_iso  = "2020-12-18T14:03:50-05:00"
  boot_time_unix = 1608318230
  host_id        = "3be84373-35bd-442d-988b-fbf92fe058fa"
  hostname       = "SATELLITE"
  kernel_version = "4.19.128-microsoft-standard"
  os_family      = "debian"
  os_name        = "ubuntu"
  os_type        = "linux"
  os_version     = "18.04"
  uptime_seconds = 980930
}

net {
  ipv4_private         = "192.168.89.154"
  ipv4_private_netmask = "/24"
  ipv4_public          = ""

  nameservers = [
    "10.0.0.1",
    "10.0.0.2",
  ]
}

