require 'formula'

class ComposeRepo < Formula
  desc "syncronize your docker-compose repos and rependancies"
  homepage "https://github.com/kron4eg/compose-repo"
  url "https://github.com/kron4eg/compose-repo/releases/download/v0.0.5/compose-repo_darwin_amd64.zip"
  version "0.0.5"
  sha256 "20aa62198429f02e7d5f7a6530c8a6efdd2af0c036d9745507d32a40d85f5603"
  depends_on :arch => :x86_64

  def install
    bin.install "compose-repo"
  end
end
