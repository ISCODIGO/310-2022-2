/** A JUnit test class for lists. */

import java.io.*;

public class ListTest
    extends junit.framework.TestCase
{

  private List<Integer> L1;
  private List<Integer> L2;
  private List<Object> L3;
/** @return True if k is in list L, false otherwise */
public static boolean find(List<Integer> L, int k) {
  for (L.moveToStart(); L.currPos()<L.length(); L.next())
    if (k == L.getValue()) return true;    // Found k
  return false;                            // k not found
}
  /**
   * This method is automatically called once before each test case method,
   * so that all the variables are cleanly initialized for each test.
   */
  public void setUp()
  {
    L1 = new LList<Integer>();
    L2 = new LList<Integer>(15);
    L3 = new LList<Object>();
  }

  public void testRemove()
  {
    L2.append(1);
    assertEquals("< | 1 >", L2.toString());
    assertEquals(1, (int)L2.remove());
    assertEquals("< | >", L2.toString());
    assertEquals(null, L2.remove());
  }

  public void testAppend()
  {
    L2.append(10);
    assertEquals("< | 10 >", L2.toString());
    L2.append(20);
    L2.append(15);
    assertEquals("< | 10 20 15 >", L2.toString());
  }

  public void testFind()
  {
    L1.moveToStart();
    L1.insert(39);
    L1.next();
    L1.insert(9);
    L1.insert(5);
    L1.append(4);
    L1.append(3);
    L1.append(2);
    L1.append(1);
    assertEquals("< 39 | 5 9 4 3 2 1 >", L1.toString());
    assertEquals(7, L1.length());

    assertEquals(true, find(L1, 3));
    assertEquals(false, find(L1, 29));
    assertEquals(true, find(L1, 5));
  }

  public void testListOfObjects()
  {
    assertEquals("< | >", L3.toString());
    L3.insert(3);
    assertEquals("< | 3 >", L3.toString());
    L3.insert("Hello");
    assertEquals("< | Hello 3 >", L3.toString());
  }

  public void testNext() {
    L2.append(10);
    L2.append(20);
    L2.append(15);
    L2.moveToStart();
    L2.next();
    assertEquals(20, (int)L2.getValue());
  }

  public void anotherTest() {
    L2.clear();
    L2.moveToStart();
    L2.insert(1);
    L2.insert(2);
    L2.moveToPos(2);
    L2.insert(3);
    L2.clear();
    assertEquals("< | >", L2.toString());
  }
}
