class Makesplash < Formula
  desc "Generate mobile app icons in all resolutions for both iOS and Android"
  homepage "https://github.com/beplus/makesplash"
  url "https://github.com/beplus/makesplash/releases/download/v0.0.4/makesplash_0.0.4_darwin_amd64.tar.gz"
  version "0.0.4"
  sha256 "95380c2b848ac3c30983d977f4566f17da16fe55d9ae7ec700c59d3753b665c6"

  def install
    bin.install "makesplash"
  end

  test do
    
  end
end
