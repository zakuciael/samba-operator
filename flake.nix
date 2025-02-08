{
  description = "An operator for a Samba as a service on PVCs in kubernetes";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=cf4f89ec3f6e7f45ca0e3a2f9f45493088952868";
    flake-parts.url = "github:hercules-ci/flake-parts";
  };

  outputs =
    {
      flake-parts,
      ...
    }@inputs:
    flake-parts.lib.mkFlake { inherit inputs; } {
      debug = true;
      systems = [
        "x86_64-linux"
      ];

      perSystem =
        {
          pkgs,
          ...
        }:
        {
          devShells.default = pkgs.mkShell {
            name = "samba-operator";
            hardeningDisable = ["all"];
            nativeBuildInputs = with pkgs; [
              gnumake
              kubernetes-controller-tools
              golangci-lint
              kustomize
              revive
              yq-go
              go
            ];
          };
        };
    };
}
