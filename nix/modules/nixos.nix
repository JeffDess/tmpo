{
  self ? null,
}:
{
  lib,
  pkgs,
  config,
  ...
}:
let
  cfg = config.programs.tmpo;
  fallbackPackage = pkgs.callPackage ../package.nix { srcPath = ../../.; };
  defaultPackage = if self == null then fallbackPackage else self.packages.${pkgs.system}.default;
in
{
  options.programs.tmpo = {
    enable = lib.mkEnableOption "tmpo CLI time tracker";

    package = lib.mkOption {
      type = lib.types.package;
      default = defaultPackage;
      description = "tmpo package to install.";
    };

    devMode = lib.mkOption {
      type = lib.types.bool;
      default = false;
      description = "Use ~/.tmpo-dev by setting TMPO_DEV=1.";
    };
  };

  config = lib.mkIf cfg.enable {
    environment.systemPackages = [ cfg.package ];
    environment.variables = lib.mkIf cfg.devMode {
      TMPO_DEV = "1";
    };
  };
}
