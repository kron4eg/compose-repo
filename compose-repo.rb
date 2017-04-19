require 'formula'

class ComposeRepo < Formula
  desc "syncronize your docker-compose repos and rependancies"
  homepage "https://github.com/kron4eg/compose-repo"
  url "https://github.com/kron4eg/compose-repo/releases/download/v0.0.6/compose-repo_darwin_amd64.zip"
  version "0.0.6"
  sha256 "ec93caeaf6b8ed64a8fcb33ac4c361ab19cfbc062ee770d578863d1e20c30ff9"
  depends_on :arch => :x86_64

  def install
    bin.install "compose-repo"
  end
end
