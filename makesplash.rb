class Makesplash < Formula
  desc "Generate mobile app icons in all resolutions for both iOS and Android"
  homepage "https://github.com/beplus/makesplash"
  url "https://github.com/beplus/makesplash/releases/download/v0.0.6/makesplash_0.0.6_darwin_amd64.tar.gz"
  version "0.0.6"
  sha256 "d1af1861226b0b4f95cb0ccd53bd40bd0da1217c56570c6607272ff8b037f7b9"

  def install
    bin.install "makesplash"
  end

  test do
    
  end
end
