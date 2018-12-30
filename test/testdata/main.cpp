// Picked from http://www.cplusplus.com/reference/algorithm/binary_search/

#include <iostream>
#include <algorithm>
#include <vector>

bool compare (int i,int j) { 
    return (i<j); 
}

/*
    Entry point
*/
int main () {
  // search target
  int myints[] = {1,2,3,4,5,4,3,2,1};
  std::vector<int> v(myints,myints+9);

  // using default comparison:
  std::sort (v.begin(), v.end());

  std::cout << "looking for a 3... ";
  if (std::binary_search (v.begin(), v.end(), 3))
    std::cout << "found!\n"; else std::cout << "not found.\n";

  // using compare as comp:
  std::sort (v.begin(), v.end(), compare);

  std::cout << "looking for a 6... ";
  if (std::binary_search (v.begin(), v.end(), 6, compare))
    std::cout << "found!\n"; else std::cout << "not found.\n";

  return 0;
}


/*
    Output:

    looking for a 3... found!
    looking for a 6... not found.
*/