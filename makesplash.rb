class Makesplash < Formula
  desc "Generate mobile app icons in all resolutions for both iOS and Android"
  homepage "https://github.com/beplus/makesplash"
  url "https://github.com/beplus/makesplash/releases/download/v0.0.1/makesplash_0.0.1_darwin_amd64.tar.gz"
  version "0.0.1"
  sha256 "883b80944a9217ab2b239997508a6d24515ce3951e5d160d5bcf25cc201cb97b"

  def install
    bin.install "makesplash"
  end

  test do
    
  end
end
