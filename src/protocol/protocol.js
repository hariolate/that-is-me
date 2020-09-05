/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/light");

var $root = ($protobuf.roots["default"] || ($protobuf.roots["default"] = new $protobuf.Root()))
.setOptions({
  go_package: ".;protocol"
})
.addJSON({
  UserAKillEvent: {
    fields: {
      objectId: {
        type: "uint32",
        id: 1
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
  UserBMoveEvent: {
    fields: {
      to: {
        type: "Position",
        id: 1
      }
    }
  },
  UserBInitEvent: {
    fields: {
      thatObjectIsMe: {
        type: "uint32",
        id: 1
      }
    }
  },
  GameStateChangeEvent: {
    fields: {
      type: {
        type: "Type",
        id: 1
      }
    },
    nested: {
      Type: {
        values: {
          Ready: 0,
          End: 1
        }
      }
    }
  },
  ClientMatchStatusRequest: {
    fields: {
      type: {
        type: "Type",
        id: 1
      }
    },
    nested: {
      Type: {
        values: {
          Ready: 0,
          End: 1
        }
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
  ChatSessionEvent: {
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
          End: 1
        }
      }
    }
  },
  UpdateObjectsLocationEvent: {
    fields: {
      objects: {
        rule: "repeated",
        type: "Object",
        id: 1
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
          Ready: 0,
          GameBegin: 1,
          GameEnd: 2,
          Exit: 3,
          EnterChat: 4,
          ExitChat: 5
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
          Cancel: 1
        }
      }
    }
  },
  MatchMakingResponse: {
    fields: {
      type: {
        type: "Type",
        id: 1
      }
    },
    nested: {
      Type: {
        values: {
          Ready: 0,
          NotAvailable: 1
        }
      }
    }
  },
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
      sender: {
        type: "Sender",
        id: 2
      },
      raw: {
        type: "RawMessage",
        id: 3
      },
      timestamp: {
        type: "google.protobuf.Timestamp",
        id: 4
      },
      timeRemaining: {
        type: "google.protobuf.Timestamp",
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
