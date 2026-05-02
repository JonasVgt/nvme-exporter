{
  description = "NVME Exporter for Prometheus";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };

        nvmeExporter = pkgs.buildGoModule {
          pname = "nvme-exporter";
          version = "0.1.0";

          src = ./.;

          # First build will fail and tell you the correct vendorHash.
          vendorHash = null;

          subPackages = [ "." ];

          meta = with pkgs.lib; {
            description = "NVME Exporter for Prometheus";
            license = licenses.mit;
            platforms = platforms.linux;
          };
        };
      in
      {
        packages.default = nvmeExporter;

        apps.default = {
          type = "app";
          program = "${nvmeExporter}/bin/nvme-exporter";
        };

        devShells.default = pkgs.mkShell {
          packages = [
            pkgs.go
            pkgs.gopls
            pkgs.gotools
          ];
        };
      });
}