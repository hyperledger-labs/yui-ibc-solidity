{
  description = "IBC Solidity";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

    ethereum = {
      inputs.nixpkgs.follows = "nixpkgs";
      url =
        "github:nix-community/ethereum.nix";
    };
    solidity-protobuf = {
      flake = false;
      url =
        "github:datachainlab/solidity-protobuf";
    };
  };

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
      ];
      systems = [ "x86_64-linux" ];
      perSystem = { config, self', inputs', pkgs, system, ... }: {
        devShells.default = pkgs.mkShell {
          SOLPB_DIR = inputs.solidity-protobuf;
          packages =
            [
              inputs.ethereum.packages.${system}.foundry
              pkgs.solc
              pkgs.go
              pkgs.protobuf
              pkgs.python3
              pkgs.python3Packages.protobuf3
              pkgs.python3Packages.wrapt
              pkgs.python3Packages.google
              pkgs.nodejs
            ];
          enterShell = '' 
             npm install
             '';
        };

      };
      flake = { };
    };
}
