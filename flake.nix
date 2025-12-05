{
  description = "advent of code monorepo";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
    systems.url = "github:nix-systems/default";
  };

  outputs = { self, nixpkgs, systems }:
    let
      supportedSystems = import systems;
      forAllSystems = f: nixpkgs.lib.genAttrs supportedSystems (system: f system);
    in
    {
      devShells = forAllSystems (system:
        let
          pkgs = import nixpkgs { inherit system; };
        in
        {
          default = pkgs.mkShell {
            packages = with pkgs; [
              # go
              go
              gopls
              gotools
              go-tools

              # c
              zig
              zls
              gdb
            ];

            GOTOOLCHAIN = "local";

            shellHook = ''
              export PATH="$PWD:$PATH"
            '';
          };
        }
      );
    };
}

