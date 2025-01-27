{
  description = "Wikipedia Golf";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

    gnumake42.url = "github:NixOS/nixpkgs/a9858885e197f984d92d7fe64e9fff6b2e488d40";
    go.url = "github:NixOS/nixpkgs/3281bec7174f679eabf584591e75979a258d8c40";
    gopls.url = "github:NixOS/nixpkgs/3281bec7174f679eabf584591e75979a258d8c40";
    typescript-language-server.url = "github:NixOS/nixpkgs/9693852a2070b398ee123a329e68f0dab5526681";
    esbuild.url = "github:NixOS/nixpkgs/3281bec7174f679eabf584591e75979a258d8c40";
    tailwindcss.url = "github:NixOS/nixpkgs/3281bec7174f679eabf584591e75979a258d8c40";
    firebase-tools.url = "github:NixOS/nixpkgs/9693852a2070b398ee123a329e68f0dab5526681";
  };

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin" ];
      perSystem = { config, self', inputs', pkgs, system, ... }: {
        packages.default = self'.packages.wgserver;
        packages.wgserver = pkgs.buildGoModule {
          name = "wgserver";
          version = "2.0.0";
          src = ./.;
          vendorHash = "sha256-GBSLR/hzPoX/G8aEK+ATm+B4NMv7SQ+ATYP/4jFBt18=";
          preBuild = ''
            cp -r ${self'.packages.wgserver-npm}/dist ./client/dist 
          '';
        };
        packages.wgserver-npm = pkgs.buildNpmPackage {
          pname = "wgserver-npm";
          version = "0.2.0";
          src = ./.;
          npmDepsHash = "sha256-76XRdNhTx5ajF+mBpBpNxCRc9lEqPBy0DCzYZbviB4A=";
          nativeBuildInputs = [
            inputs'.tailwindcss.legacyPackages.tailwindcss
          ];
          postBuild = ''
            mkdir -p $out
            cp -r ./client/dist $out/dist
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
            inputs'.typescript-language-server.legacyPackages.nodePackages.typescript-language-server
            inputs'.esbuild.legacyPackages.esbuild
            inputs'.tailwindcss.legacyPackages.tailwindcss
            inputs'.firebase-tools.legacyPackages.firebase-tools
            inputs'.tailwindcss.legacyPackages.tailwindcss-language-server
          ];
        };
      };
    };
}
