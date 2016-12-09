SRC=git@github.com:kiwamunet/zopfli.git
VENDOR=vendor
ZOPFLI=zopfli
LIB=libzopflipng.a
RM=rm
GIT=git
HEADER=zopflipng_lib.h
MAKE=make
MAKEFLAG=zopflipng
CP=cp -pf

all:
	$(RM) -rf $(ZOPFLI)
	$(GIT) clone $(SRC) 
	$(MAKE) $(MAKEFLAG) -C $(ZOPFLI)/
	$(CP) $(ZOPFLI)/$(LIB) $(VENDOR)/$(LIB)
	$(CP) $(ZOPFLI)/src/zopflipng/$(HEADER) $(VENDOR)/$(HEADER)
	$(RM) -rf $(ZOPFLI)
