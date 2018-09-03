class Clif < Formula
    desc "Bookmarks for cli commands"
    homepage "https://github.com/dutchakdev/clif"
    url "https://github.com/dutchakdev/clif/releases/download/v0.0.1/clif_0.0.1_Darwin_x86_64.tar.gz"
    sha256 "78599eb003d812784dfe20230c9e6b87deb2d028d06dcafabdc50cfcd82a142e"

    def install
        bin.install "clif"
    end

    test do
      clif -v
    end
end