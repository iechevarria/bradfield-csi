#include "vec.h"


// changes (all improved run time):
// - move length out of loop check
// - don't do bounds checks
// - do some loop unrolling

data_t dotproduct(vec_ptr u, vec_ptr v) {
   data_t sum1 = 0, sum2 = 0, sum3 = 0, sum4 = 0;
   long i = 0;

   int len = vec_length(u);

   // we can assume both vectors are same length
   for (; i < len-3; i+=4) {
        sum1 += u->data[i] * v->data[i];
        sum2 += u->data[i+1] * v->data[i+1];
        sum3 += u->data[i+2] * v->data[i+2];
        sum4 += u->data[i+3] * v->data[i+3];
   }

   // handle last few elements
   for (; i < len; i++) {
        sum1 += u->data[i] * v->data[i];
   }

   return sum1 + sum2 + sum3 + sum4;
}
