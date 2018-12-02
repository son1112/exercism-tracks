(defpackage #:hello-world
  (:use #:common-lisp)
  (:export #:hello-world)
  (:nicknames #:hw))

(in-package #:hello-world)

(defun hellow-world
  "Return the string: Hello, World!"
  (format t "Hello, World!")
  )
