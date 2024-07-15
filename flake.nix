{
  description = "Shoppinglist server";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    let
      # Application metadata
      name = "shoppinglist-server";

      image = {
        inherit name;
        registry = "ghcr.io";
        owner = "kmjayadeep";
      };
    in

    flake-utils.lib.eachDefaultSystem
      (system:
        let
          pkgs = import nixpkgs {
            inherit system;
          };
        in
        {
          packages = rec {
            # Enables us to build the Go service by running plain old `nix build`
            default = shoppinglist-server;

            shoppinglist-server = pkgs.buildGo122Module {
              inherit name;
              vendorHash = null;
              src = ./.;
            };

            # Intended only for CI. The image fails to run if you build it on a
            # non-`x86_64-linux` system, despite the build succeeding. There are
            # ways around this in Nix but in this case we only need to build the
            # image in CI.
            docker =
              pkgs.dockerTools.buildImage {
                name = "${image.registry}/${image.owner}/${image.name}";
                tag  = "latest";
                config = {
                  Cmd = [ "${shoppinglist-server}/bin/${name}" ];
                  ExposedPorts."8080/tcp" = { };
                };
              };
          };

          # Cross-platform development environment (including CI)
          devShells.default = pkgs.mkShell {
            # Packages available in the environment
            packages = with pkgs; [
              # Golang
              go_1_22
              gotools # goimports, etc.

              # Utilities
              httpie # For arbitrary HTTP calls
              go-swag

              # DevOps
              kubectl # Kubernetes CLI
              kubectx # Kubernetes context management utility
            ];

            shellHook = ''
              exec zsh
            '';
          };
        });
}
