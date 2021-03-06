double **matrix_alloc(int m, int n);
void matrix_free(double **matrix, int m);
void matrix_multiply(double **c, double **a, double **b, int a_height, int b_height, int b_width);
void fast_matrix_multiply(double **c, double **a, double **b, int a_height, int b_height, int b_width);
