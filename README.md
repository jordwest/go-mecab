GoLang bindings for MeCab
=========================

Installation
------------
Install MeCab from source as usual.

First download the source and dictionary archives from http://mecab.googlecode.com/ then:

Install MeCab

    tar zxfv mecab-X.X.tar.gz
    cd mecab-X.X
    ./configure --enable-utf8-only
    make
    make check
    sudo make install

Reload libraries

    sudo ldconfig

Install the dictionary

    tar zxfv mecab-ipadic-X.tar.gz
    cd mecab-ipadic-X
    ./configure --with-charset=utf8
    make
    sudo make install


