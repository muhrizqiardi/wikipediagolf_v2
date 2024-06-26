{
  description = "Wikipedia Golf";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

    gnumake42.url = "github:NixOS/nixpkgs/a9858885e197f984d92d7fe64e9fff6b2e488d40";
    go.url = "github:NixOS/nixpkgs/3281bec7174f679eabf584591e75979a258d8c40";
    gopls.url = "github:NixOS/nixpkgs/3281bec7174f679eabf584591e75979a258d8c40";
    typescript-language-server.url = "github:NixOS/nixpkgs/9693852a2070b398ee123a329e68f0dab5526681";
  };

  outputs = inputs@{ flake-parts, templ, go, gopls, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin" ];
      perSystem = { config, self', inputs', pkgs, system, ... }: {
        packages.default = self'.packages.wgserver;
        packages.wgserver = pkgs.buildGoModule {
          name = "wgserver";
          version = "2.0.0";
          subPackages = [ "cmd/wgserver" ];
          src = ./server;
          vendorHash = "sha256-1wXTcB8QfeqWS9DUJtxPykSnMy2YeTRpkoZ10WJT6ig=";
          preBuild = ''
            cp -r ${self'.packages.wgserver-npm}/dist ./dist
          '';
        };
        packages.wgserver-docker = pkgs.dockerTools.buildImage {
          name = "muhrizqiardi/wgserver";
          tag = "2.0.0";

          copyToRoot = pkgs.buildEnv {
            name = "image-root";
            pathsToLink = [ "/bin" "/etc" "/var" ];
            paths = [
              pkgs.dockerTools.fakeNss
              pkgs.dockerTools.binSh
              pkgs.dockerTools.usrBinEnv
              pkgs.dockerTools.caCertificates
              self'.packages.wgserver
            ];
          };
          config = {
            Cmd = [
              "${self'.packages.wgserver}/bin/wgserver"
            ];
            ExposedPorts = {
              "3000/tcp" = { };
            };
          };
        };
        devShells.default = pkgs.mkShell {
          nativeBuildInputs = [
            inputs'.gnumake42.legacyPackages.gnumake42
            inputs'.go.legacyPackages.go
            inputs'.gopls.legacyPackages.gopls
            inputs'.typescript-language-server.lecacyPackages.typescript-language-server
          ];
        };
      };
    };
}
