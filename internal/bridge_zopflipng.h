#include "internal/zopfli/src/zopflipng/zopflipng_lib.h"

void setfilters(CZopfliPNGOptions* png_options, int size, int* filters);
void setchunks(CZopfliPNGOptions* png_options, int size, char** chunks);
void freeopts(CZopfliPNGOptions* png_options);

