<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hello.proto

namespace GPBMetadata;

class Hello
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(hex2bin(
            "0a9a010a0b68656c6c6f2e70726f746f120568656c6c6f221c0a0c48656c" .
            "6c6f52657175657374120c0a046e616d65180120012809221d0a0a48656c" .
            "6c6f5265706c79120f0a076d657373616765180120012809323f0a074772" .
            "656574657212340a0853617948656c6c6f12132e68656c6c6f2e48656c6c" .
            "6f526571756573741a112e68656c6c6f2e48656c6c6f5265706c79220062" .
            "0670726f746f33"
        ), true);

        static::$is_initialized = true;
    }
}

