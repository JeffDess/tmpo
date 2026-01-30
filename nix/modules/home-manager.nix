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
  settingsFile = lib.filterAttrs (_: value: value != null) {
    inherit (cfg.settings) currency timezone;
    date_format = cfg.settings.dateFormat;
    time_format = cfg.settings.timeFormat;
    export_path = cfg.settings.exportPath;
  };
  hasSettings = builtins.length (builtins.attrNames settingsFile) > 0;
  yamlSettings = lib.generators.toKeyValue {
    mkKeyValue = lib.generators.mkKeyValueDefault { } ": ";
  } settingsFile;
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

    settings = {
      currency = lib.mkOption {
        type = lib.types.nullOr lib.types.str;
        default = null;
        description = "ISO 4217 currency code (ex: USD).";
      };

      dateFormat = lib.mkOption {
        type = lib.types.nullOr lib.types.str;
        default = null;
        description = "Date format (MM/DD/YYYY, DD/MM/YYYY, YYYY-MM-DD).";
      };

      timeFormat = lib.mkOption {
        type = lib.types.nullOr lib.types.str;
        default = null;
        description = "Time format (24-hour, 12-hour (AM/PM)).";
      };

      timezone = lib.mkOption {
        type = lib.types.nullOr lib.types.str;
        default = null;
        description = "IANA timezone (ex: America/New_York).";
      };

      exportPath = lib.mkOption {
        type = lib.types.nullOr lib.types.str;
        default = null;
        description = "Default export path for csv/json exports.";
      };
    };
  };

  config = lib.mkIf cfg.enable {
    home = {
      packages = [ cfg.package ];
      sessionVariables = lib.mkIf cfg.devMode {
        TMPO_DEV = "1";
      };

      file = lib.mkIf hasSettings {
        ".tmpo/config.yaml".text = yamlSettings;
      };
    };
  };
}
