
fn pow (base, exp) {
    result = base;
    for (i = 0; i<exp; i = i + 1;) {
        result = result * base;
        print("result:",result);
    }
    return result;
}

print(pow(2, 3));

