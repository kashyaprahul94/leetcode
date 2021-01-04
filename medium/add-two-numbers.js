function ListNode(val, next) {
  this.val = val === undefined ? 0 : val;
  this.next = next === undefined ? null : next;
}

ListNode.prototype.toString = function () {
  let head = this;
  let result = `Start`;

  while(head) {
    result = `${result} -> ${head.val}`;
    head = head.next;
  }

  result = `${result} -> End`;

  return result;
}

ListNode.prototype.toArray = function () {
  let head = this;
  const result = [];

  while(head) {
    result.push(head.val);
    head = head.next;
  }

  return result;
}

/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {
  let resultList = new ListNode(0);
  let current = resultList;

  let carry = 0;

  while (l1 || l2) {
    let sum = (l1?.val ?? 0) + (l2?.val ?? 0) + carry;

    if (sum >= 10) {
      sum = sum - 10;
      carry = 1;
    } else {
      carry = 0;
    }

    current.next = new ListNode(sum);
    current = current.next;

    if (l1) {
      l1 = l1.next;
    }
    if (l2) {
      l2 = l2.next;
    }
  }

  if (carry > 0) {
    current.next = new ListNode(carry);
  }

  return resultList.next;
};

let l1, l2;

l1 = new ListNode(2, new ListNode(4, new ListNode(3)));
l2 = new ListNode(5, new ListNode(6, new ListNode(4)));

console.log(addTwoNumbers(l1, l2).toArray());

l1 = new ListNode(9,
  new ListNode(9,
    new ListNode(9,
      new ListNode(9,
        new ListNode(9,
          new ListNode(9,
            new ListNode(9)
          )
        )
      )
    )
  )
);
l2 = new ListNode(9,
  new ListNode(9,
    new ListNode(9,
      new ListNode(9)
    )
  )
);

console.log(addTwoNumbers(l1, l2).toArray());

l1 = new ListNode(1, new ListNode(1));
l2 = new ListNode(9, new ListNode(9, new ListNode(9)));

console.log(addTwoNumbers(l1, l2).toArray());

l1 = new ListNode(0, new ListNode(0, new ListNode(1)));
l2 = new ListNode(9);

console.log(addTwoNumbers(l1, l2).toArray());
