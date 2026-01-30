{
  description = "tmpo CLI time tracker";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      ...
    }:
    let
      lib = nixpkgs.lib;
      sourceInfo = self.sourceInfo or { };
      ref = sourceInfo.ref or "";
      version = if lib.hasPrefix "v" ref then lib.removePrefix "v" ref else "unstable";
    in
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
        goPkg = pkgs.go_1_25 or pkgs.go;
        rev = self.rev or (self.dirtyRev or "dirty");
        lastModified = sourceInfo.lastModifiedDate or self.lastModifiedDate or "19700101000000";
        date =
          "${builtins.substring 0 4 lastModified}-"
          + "${builtins.substring 4 2 lastModified}-"
          + "${builtins.substring 6 2 lastModified}T"
          + "${builtins.substring 8 2 lastModified}:"
          + "${builtins.substring 10 2 lastModified}:"
          + "${builtins.substring 12 2 lastModified}Z";
        package = pkgs.callPackage ./nix/package.nix {
          inherit version;
          commit = rev;
          inherit date;
          go = goPkg;
          srcPath = ./.;
        };
      in
      {
        packages = {
          default = package;
          tmpo = package;
        };
        apps.default = flake-utils.lib.mkApp { drv = package; };
        devShells.default = pkgs.mkShell {
          packages = [
            goPkg
            pkgs.gopls
            pkgs.gotools
          ];
        };
      }
    )
    // {
      nixosModules.tmpo = import ./nix/modules/nixos.nix { inherit self; };
      homeManagerModules.tmpo = import ./nix/modules/home-manager.nix { inherit self; };
    };
}
