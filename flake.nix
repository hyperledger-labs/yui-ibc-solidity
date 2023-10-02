{
  description = "IBC Solidity";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

    ethereum = {
      inputs.nixpkgs.follows = "nixpkgs";
      url =
        "github:nix-community/ethereum.nix";
    };
  };

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
      ];
      systems = [ "x86_64-linux" ];
      perSystem = { config, self', inputs', pkgs, system, ... }: {
        devShells.default = pkgs.mkShell {
          packages =
            [
              inputs.ethereum.packages.${system}.foundry
              pkgs.solc
              pkgs.go
              pkgs.protobuf
            ];
        };

      };
      flake = { };
    };
}
