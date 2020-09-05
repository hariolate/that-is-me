/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/light");

var $root = ($protobuf.roots["default"] || ($protobuf.roots["default"] = new $protobuf.Root()))
.setOptions({
  go_package: ".;protocol"
})
.addJSON({
  RawMessage: {
    fields: {
      message: {
        type: "string",
        id: 1
      }
    }
  },
  Message: {
    fields: {
      uid: {
        type: "uint32",
        id: 1
      },
      raw: {
        type: "RawMessage",
        id: 2
      },
      timestamp: {
        type: "google.protobuf.Timestamp",
        id: 3
      }
    }
  },
  google: {
    nested: {
      protobuf: {
        nested: {
          Timestamp: {
            fields: {
              seconds: {
                type: "int64",
                id: 1
              },
              nanos: {
                type: "int32",
                id: 2
              }
            }
          }
        }
      }
    }
  }
});

module.exports = $root;
