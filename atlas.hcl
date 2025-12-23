data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/domain",
    "--dialect", "postgres", // | postgres | sqlite | sqlserver
  ]
}
env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/18/dev"
  url = "postgres://root@localhost:5432/jagoanbunda?&sslmode=disable"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
