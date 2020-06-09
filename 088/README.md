题目：[合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array)

本题采用新建一个数组 `nums3`

同时循环 `nums1` 和 `nums2`，并判断两数组当前的循环值大小

将小的值插入 `nums3` 中，最后循环赋值给 `nums1`

这里需要注意一个事，`golang` 虽然都是值传递，但是在 StackOverFlow 上有人做出了解释， [Are golang slices pass by value?](https://link.jianshu.com/?t=http://stackoverflow.com/questions/39993688/are-golang-slices-pass-by-value) 摘要如下：

> Everything in Go is passed by value. Slices too. But a slice value is a header, describing a contiguous section of a backing array, and a slice value only contains a pointer to the array where the elements are actually stored. The slice value does not include its elements (unlike arrays).
>
> So when you pass a slice to a function, a copy will be made from this header, including the pointer, which will point to the same backing array. Modifying the elements of the slice implies modifying the elements of the backing array, and so all slices which share the same backing array will "observe" the change.

大概意思就是：slice 作为参数传递给函数其实是传递 slice 的值，这个值被称作一个 `header` ，**它只包含了一个指向底层数组的指针**。当向函数传递一个 slice ，将复制一个 header 的副本，这个副本包含一个指向同一个底层数组的指针。修改 slice 的元素间接地修改底层数组的元素，也就是所有指向同一个底层数组的 slice 会响应这个变化，主函数的 slice 也就一同修改了 s[0] 的值。

所以最后直接赋值 `nums1 = nums3` 是不可取的，会改变内部的地址

而循环赋值则改变的是当前地址的值