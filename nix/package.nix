{
  lib,
  buildGoModule,
  go,
  srcPath ? lib.cleanSource ../.,
  version ? "0.4.5",
  commit ? "dirty",
  date ? "1970-01-01T00:00:00Z",
}:
buildGoModule {
  pname = "tmpo";
  inherit version go;
  src = srcPath;

  vendorHash = "sha256-4S5ESwFJ8JYKugNJgTkxUywZM7dTqDlDIvXcNRot5b4=";
  subPackages = [ "." ];

  ldflags = [
    "-s"
    "-w"
    "-X github.com/DylanDevelops/tmpo/cmd/utilities.Version=${version}"
    "-X github.com/DylanDevelops/tmpo/cmd/utilities.Commit=${commit}"
    "-X github.com/DylanDevelops/tmpo/cmd/utilities.Date=${date}"
  ];

  env = {
    CGO_ENABLED = "0";
  };

  doCheck = true;

  meta = with lib; {
    description = "Minimal CLI time tracker for developers";
    homepage = "https://github.com/DylanDevelops/tmpo";
    license = licenses.mit;
    mainProgram = "tmpo";
  };
}
