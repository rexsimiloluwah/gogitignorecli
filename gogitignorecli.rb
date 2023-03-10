# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Gogitignorecli < Formula
  desc ""
  homepage "https://github.com/rexsimiloluwah/gogitignorecli"
  version "0.1.0"

  on_macos do
    url "https://github.com/rexsimiloluwah/gogitignorecli/releases/download/v0.1.0/gogitignorecli_0.1.0_darwin_all.tar.gz"
    sha256 "ce30b08f99332b8618c7ff124e8e5d7623f1bd1edb85520e8eb6c49d79842c10"

    def install
      bin.install "gogitignorecli"
    end
  end

  on_linux do
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/rexsimiloluwah/gogitignorecli/releases/download/v0.1.0/gogitignorecli_0.1.0_linux_armv6.tar.gz"
      sha256 "117a8aa3b307975e0d84e7b67926d4c2d72d07baedbd349db49962ff7c68c6f7"

      def install
        bin.install "gogitignore"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/rexsimiloluwah/gogitignorecli/releases/download/v0.1.0/gogitignorecli_0.1.0_linux_amd64.tar.gz"
      sha256 "7de884051ff6ac2a339115b306a73a601bc7cb5c39547d1212774c9d3c3efdf1"

      def install
        bin.install "gogitignore"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/rexsimiloluwah/gogitignorecli/releases/download/v0.1.0/gogitignorecli_0.1.0_linux_arm64.tar.gz"
      sha256 "bfed7f8800424afb0b5b8ff8d45aa0a20160b36c510493cf9a0a0e290d2fb82b"

      def install
        bin.install "gogitignore"
      end
    end
  end
end
