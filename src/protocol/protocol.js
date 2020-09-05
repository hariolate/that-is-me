/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/light");

var $root = ($protobuf.roots["default"] || ($protobuf.roots["default"] = new $protobuf.Root()))
.setOptions({
  go_package: ".;protocol"
})
.addJSON({
  RawChatMessageEvent: {
    fields: {
      message: {
        type: "string",
        id: 1
      }
    }
  },
  NewMessageEvent: {
    fields: {
      sender: {
        type: "Sender",
        id: 2
      },
      raw: {
        type: "RawChatMessageEvent",
        id: 3
      },
      timestamp: {
        type: "google.protobuf.Timestamp",
        id: 4
      },
      timeRemaining: {
        type: "int64",
        id: 5
      }
    },
    nested: {
      Sender: {
        values: {
          A: 0,
          B: 1
        }
      }
    }
  },
  Position: {
    fields: {
      x: {
        type: "uint32",
        id: 1
      },
      y: {
        type: "uint32",
        id: 2
      }
    }
  },
  Object: {
    fields: {
      position: {
        type: "Position",
        id: 1
      },
      objectId: {
        type: "uint32",
        id: 2
      }
    }
  },
  UserAKillEvent: {
    fields: {
      objectId: {
        type: "uint32",
        id: 1
      }
    }
  },
  UserBMoveEvent: {
    fields: {
      to: {
        type: "Position",
        id: 1
      }
    }
  },
  UserAInitEvent: {
    fields: {
      objects: {
        rule: "repeated",
        type: "Object",
        id: 1
      }
    }
  },
  UserBInitEvent: {
    fields: {
      objects: {
        rule: "repeated",
        type: "Object",
        id: 1
      },
      thatObjectIsMe: {
        type: "uint32",
        id: 2
      }
    }
  },
  ClientReadyEvent: {
    fields: {}
  },
  ObjectsLocationUpdateEvent: {
    fields: {
      objects: {
        rule: "repeated",
        type: "Object",
        id: 1
      }
    }
  },
  GameResultEvent: {
    fields: {
      type: {
        type: "Type",
        id: 1
      }
    },
    nested: {
      Type: {
        values: {
          Success: 0,
          Failure: 1
        }
      }
    }
  },
  MatchMakingRequest: {
    fields: {
      type: {
        type: "Type",
        id: 1
      }
    },
    nested: {
      Type: {
        values: {
          Begin: 0,
          Cancel: 1,
          Accept: 2,
          NotAccept: 3
        }
      }
    }
  },
  MatchMakingResponse: {
    fields: {
      type: {
        type: "Type",
        id: 1
      },
      role: {
        type: "Role",
        id: 2
      }
    },
    nested: {
      Type: {
        values: {
          Available: 0,
          NotAvailable: 1
        }
      },
      Role: {
        values: {
          A: 0,
          B: 1
        }
      }
    }
  },
  MatchStateChangeEvent: {
    fields: {
      type: {
        type: "Type",
        id: 1
      }
    },
    nested: {
      Type: {
        values: {
          MatchPending: 0,
          MatchAccepted: 1,
          GameReady: 2,
          GameBegin: 3,
          GameEnd: 4,
          ChatPending: 5,
          ChatReady: 6,
          ChatBegin: 7,
          ChatEnd: 8,
          MatchExit: 9
        }
      }
    }
  },
  Wrapper: {
    fields: {
      type: {
        type: "MessageType",
        id: 1
      },
      message: {
        type: "google.protobuf.Any",
        id: 2
      }
    },
    nested: {
      MessageType: {
        values: {
          MatchMakingRequest: 0,
          MatchMakingResponse: 1,
          MatchStateChangeEvent: 2,
          UserAKillEvent: 3,
          UserBMoveEvent: 4,
          UserAInitEvent: 5,
          UserBInitEvent: 6,
          ClientReadyEvent: 7,
          ObjectsLocationUpdateEvent: 8,
          GameResultEvent: 9,
          RawChatMessageEvent: 10,
          NewMessageEvent: 11
        }
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
          },
          Any: {
            fields: {
              type_url: {
                type: "string",
                id: 1
              },
              value: {
                type: "bytes",
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
