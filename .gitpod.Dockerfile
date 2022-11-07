FROM gitpod/workspace-full

# Install LaTeX
RUN sudo apt-get -q update && \
    sudo apt-get install -yq texlive-latex-extra latexmk && \ 
    sudo rm -rf /var/lib/apt/lists/*

