package spirit

import (
	"github.com/gogap/errors"
)

var (
	ErrReceiverURNNotExist          = errors.New("receiver urn not exist")
	ErrReceiverNameDuplicate        = errors.New("receiver name duplicate")
	ErrReceiverAlreadyRunning       = errors.New("receiver already running")
	ErrReceiverDidNotHaveTranslator = errors.New("receiver did not have translator")
	ErrReceiverDidNotRunning        = errors.New("receiver did not running")
	ErrReceiverReceiveTimeout       = errors.New("receiver receive timeout")
	ErrReceiverDeliveryPutterIsNil  = errors.New("receiver delivery putter is nil")
	ErrReceiverDidNotHaveReaderPool = errors.New("receiver did not have reader pool")
	ErrReceiverTypeNotSupport       = errors.New("receiver type not support")

	ErrSenderURNNotExist          = errors.New("sender urn not exist")
	ErrSenderNameDuplicate        = errors.New("sender name duplicate")
	ErrSenderAlreadyRunning       = errors.New("sender already running")
	ErrSenderCanNotCreaterWriter  = errors.New("sender can not create writer")
	ErrSenderDidNotHaveTranslator = errors.New("sender did not have translator")
	ErrSenderDidNotRunning        = errors.New("sender did not running")
	ErrSenderDeliveryGetterIsNil  = errors.New("sender delivery getter is nil")
	ErrSenderDidNotHaveWriterPool = errors.New("sender did not have writer pool")
	ErrSenderTypeNotSupport       = errors.New("sender type not support")

	ErrRouterURNNotExist                     = errors.New("router urn not exist")
	ErrRouterNameDuplicate                   = errors.New("router name duplicate")
	ErrRouterAlreadyRunning                  = errors.New("router already running")
	ErrRouterDidNotRunning                   = errors.New("router did not running")
	ErrRouterAlreadyHaveThisInbox            = errors.New("router already added this inbox")
	ErrRouterAlreadyHaveThisOutbox           = errors.New("router already added this outbox")
	ErrRouterAlreadyHaveThisComponent        = errors.New("router already added this component")
	ErrRouterToComponentHandlerFailed        = errors.New("could not router to component handler")
	ErrRouterToComponentFailed               = errors.New("could not router to component")
	ErrRouterComponentNotExist               = errors.New("router component not exist")
	ErrRouterDidNotHaveComponentLabelMatcher = errors.New("could not router to component")
	ErrRouterHandlerCountNotEqualURNsCount   = errors.New("handler's count is not equal delivery urn's count, some component may not add")
	ErrRouterDeliveryURNFormatError          = errors.New("delivery urn format error")
	ErrRouterOnlyOneGlobalComponentAllowed   = errors.New("only one same urn global component allowed")

	ErrComponentURNNotExist       = errors.New("component urn not exist")
	ErrComponentNameDuplicate     = errors.New("component name duplicate")
	ErrComponentNameIsEmpty       = errors.New("component name is empty")
	ErrComponentURNIsEmpty        = errors.New("component URN is empty")
	ErrComponentNotExist          = errors.New("component not exist")
	ErrComponentHandlerNotExit    = errors.New("component handler not exist")
	ErrComponentDidNotHaveHandler = errors.New("component did not have handler")

	ErrDeliveryURNIsEmpty = errors.New("delivery urn is empty")

	ErrReaderURNNotExist   = errors.New("reader urn not exist")
	ErrReaderNameDuplicate = errors.New("reader name duplicate")

	ErrWriterURNNotExist   = errors.New("writer urn not exist")
	ErrWriterNameDuplicate = errors.New("writer name duplicate")

	ErrWriterIsNil              = errors.New("writer is nil")
	ErrWriterInUse              = errors.New("writer already in use")
	ErrWriterPoolAlreadyClosed  = errors.New("writer pool already closed")
	ErrWriterPoolNameDuplicate  = errors.New("writer pool name duplicate")
	ErrWriterPoolURNNotExist    = errors.New("writer pool urn not exist")
	ErrWriterPoolTooManyWriters = errors.New("writer pool have too many writers")

	ErrReaderIsNil              = errors.New("reader is nil")
	ErrReaderInUse              = errors.New("reader already in use")
	ErrReaderPoolAlreadyClosed  = errors.New("reader pool already closed")
	ErrReaderPoolTooManyReaders = errors.New("reader pool have too many readers")
	ErrReaderPoolNameDuplicate  = errors.New("reader pool name duplicate")
	ErrReaderPoolURNNotExist    = errors.New("reader pool urn not exist")

	ErrInboxURNNotExist   = errors.New("inbox urn not exist")
	ErrInboxNameDuplicate = errors.New("inbox name duplicate")

	ErrOutboxURNNotExist   = errors.New("outbox urn not exist")
	ErrOutboxNameDuplicate = errors.New("outbox name duplicate")

	ErrLabelMatcherURNNotExist   = errors.New("label matcher urn not exist")
	ErrLabelMatcherNameDuplicate = errors.New("label matcher name duplicate")

	ErrURNRewriterURNNotExist   = errors.New("urn rewriter urn not exist")
	ErrURNRewriterNameDuplicate = errors.New("urn rewriter name duplicate")

	ErrInputTranslatorURNNotExist   = errors.New("input translator urn not exist")
	ErrInputTranslatorNameDuplicate = errors.New("input translator name duplicate")

	ErrOutputTranslatorURNNotExist   = errors.New("output translator urn not exist")
	ErrOutputTranslatorNameDuplicate = errors.New("output translator name duplicate")

	ErrSpiritAlreadyRunning   = errors.New("spirit already running")
	ErrSpiritAlreadyStopped   = errors.New("spirit already stopped")
	ErrSpiritAlreadyBuilt     = errors.New("spirit already built")
	ErrSpiritNotBuild         = errors.New("spirit not build")
	ErrSpiritActorURNNotExist = errors.New("spirit actor urn not exist")

	ErrActorRouterNotExist           = errors.New("actor:router not exist")
	ErrActorComponentNotExist        = errors.New("actor:component not exist")
	ErrActorInBoxNotExist            = errors.New("actor:inbox not exist")
	ErrActorReceiverNotExist         = errors.New("actor:receiver not exist")
	ErrActorInputTranslatorNotExist  = errors.New("actor:input_translator not exist")
	ErrActorOutputTranslatorNotExist = errors.New("actor:output_translator not exist")
	ErrActorReaderPoolNotExist       = errors.New("actor:reader_pool not exist")
	ErrActorWriterNotExist           = errors.New("actor:writer not exist")
	ErrActorOutboxNotExist           = errors.New("actor:outbox not exist")
	ErrActorLabelMatcerNotExist      = errors.New("actor:label_matcher not exist")
	ErrActorSenderNotExist           = errors.New("actor:sender not exist")
	ErrActorURNRewriterNotExist      = errors.New("actor:urn_rewriter not exist")

	ErrComposeNameIsEmpty = errors.New("spirit compose name is empty")
)

const SPIRIT_ERR_NS = "SPIRIT"

var (
	ERR_MESSAGE_GRAPH_ADDRESS_NOT_FOUND = errors.TN(SPIRIT_ERR_NS, 1, "could not found next address in graph, current component name: {{.compName}}, port name: {{.portName}}")
	ERR_MESSAGE_GRAPH_IS_NIL            = errors.TN(SPIRIT_ERR_NS, 2, "component graph is nil")
	ERR_MESSAGE_ADDRESS_IS_EMPTY        = errors.TN(SPIRIT_ERR_NS, 3, "component message address is empty")
	ERR_MESSAGE_SERIALIZE_FAILED        = errors.TN(SPIRIT_ERR_NS, 4, "component message serialize failed, err: {{.err}}")

	ERR_COMPONENT_HANDLER_RETURN_ERROR = errors.TN(SPIRIT_ERR_NS, 5, "component handler return error, component name: {{.name}}, port: {{.port}}, error: {{.err}}")
	ERR_COMPONENT_HANDLER_NOT_EXIST    = errors.TN(SPIRIT_ERR_NS, 6, "component handler not exist, component name: {{.name}},  handler name: {{.handlerName}}")

	ERR_SENDER_CREDENTIAL_IS_NIL = errors.TN(SPIRIT_ERR_NS, 8, "credential is nil, type: {{.type}}, url: {{.url}}")
	ERR_SENDER_MNS_CLIENT_IS_NIL = errors.TN(SPIRIT_ERR_NS, 9, "sender mns client is nil, type: {{.type}}, url: {{.url}}")
	ERR_SENDER_SEND_FAILED       = errors.TN(SPIRIT_ERR_NS, 10, "component message send failed, type: {{.type}}, url: {{.url}}, err:{{.err}}")

	ERR_SENDER_CREATE_FAILED    = errors.TN(SPIRIT_ERR_NS, 11, "create sender failed, urn type: {{.type}}")
	ERR_SENDER_DRIVER_NOT_EXIST = errors.TN(SPIRIT_ERR_NS, 12, "message sender urn not exist, type: {{.type}}")
	ERR_SENDER_BAD_DRIVER       = errors.TN(SPIRIT_ERR_NS, 13, "bad message sender urn of {{.type}}")

	ERR_RECEIVER_CREDENTIAL_IS_NIL     = errors.TN(SPIRIT_ERR_NS, 14, "credential is nil, type: {{.type}}")
	ERR_RECEIVER_MNS_CLIENT_IS_NIL     = errors.TN(SPIRIT_ERR_NS, 15, "receiver mns client is nil, type: {{.type}}, url: {{.url}}")
	ERR_RECEIVER_DELETE_MSG_ERROR      = errors.TN(SPIRIT_ERR_NS, 16, "delete message error, type: {{.type}}, url: {{.url}}, err: {{.err}}")
	ERR_RECEIVER_UNMARSHAL_MSG_FAILED  = errors.TN(SPIRIT_ERR_NS, 17, "receiver unmarshal message failed, type: {{.type}}, err: {{.err}}")
	ERR_RECEIVER_RECV_ERROR            = errors.TN(SPIRIT_ERR_NS, 18, "receiver recv error, type: {{.type}}, url: {{.url}}, err: {{.err}}")
	ERR_RECEIVER_BAD_DRIVER            = errors.TN(SPIRIT_ERR_NS, 19, "bad message receiver urn of {{.type}}")
	ERR_RECEIVER_CREATE_FAILED         = errors.TN(SPIRIT_ERR_NS, 20, "create receiver failed, urn type: {{.type}}, url: {{.url}}")
	ERR_RECEIVER_DRIVER_NOT_EXIST      = errors.TN(SPIRIT_ERR_NS, 21, "message receiver urn not exist, type: {{.type}}")
	ERR_RECEIVER_DECODE_MESSAGE_FAILED = errors.TN(SPIRIT_ERR_NS, 22, "message receiver decode message failed, type: {{.type}}, url: {{.url}}, err: {{.err}}")

	ERR_PORT_NAME_IS_EMPTY    = errors.TN(SPIRIT_ERR_NS, 23, "port name is empty, component: {{.name}}")
	ERR_HANDLER_NAME_IS_EMPTY = errors.TN(SPIRIT_ERR_NS, 24, "handler name is empty, component: {{.name}}")

	ERR_UUID_GENERATE_FAILED = errors.TN(SPIRIT_ERR_NS, 26, "uuid generate failed, err: {{.err}}")

	ERR_PAYLOAD_GO_BAD           = errors.TN(SPIRIT_ERR_NS, 27, "payload go bad, err: {{.err}}")
	ERR_PAYLOAD_SERIALIZE_FAILED = errors.TN(SPIRIT_ERR_NS, 28, "payload serialize failed, err: {{.err}}")

	ERR_COMPONENT_HANDLER_PANIC = errors.TN(SPIRIT_ERR_NS, 29, "component handler panic,component: {{.name}}, err: {{.err}}")

	ERR_HEARTBEAT_OPTIONS_IS_NIL = errors.TN(SPIRIT_ERR_NS, 30, "heartbeat options is nil, heartbeat name: {{.name}}")

	ERR_READE_FILE_ERROR     = errors.TN(SPIRIT_ERR_NS, 31, "read file error,file: {{.file}} err: {{.err}}")
	ERR_UNMARSHAL_DATA_ERROR = errors.TN(SPIRIT_ERR_NS, 32, "unmarshal data error, err: {{.err}}")

	ERR_HEARTBEAT_ALI_JIANKONG_UID_NOT_EXIST = errors.TN(SPIRIT_ERR_NS, 33, "heartbeat of ali jiankong's uid does not exist")

	ERR_HOOK_CREATE_FAILED             = errors.TN(SPIRIT_ERR_NS, 34, "create receiver failed, urn type: {{.type}}, name: {{.name}}")
	ERR_HOOK_DRIVER_NOT_EXIST          = errors.TN(SPIRIT_ERR_NS, 35, "message hook urn not exist, type: {{.type}}")
	ERR_HOOK_BAD_DRIVER                = errors.TN(SPIRIT_ERR_NS, 36, "bad message hook urn of {{.type}}")
	ERR_MESSAGE_HOOK_ERROR             = errors.TN(SPIRIT_ERR_NS, 37, "hook message error, component: {{.name}}, port: {{.port}}, hookName: {{.hookName}}, event: {{.event}}, index: {{.index}}/{{.count}}")
	ERR_HOOK_INSTANCE_ALREADY_INITALED = errors.TN(SPIRIT_ERR_NS, 38, "hook instance already initaled, urn type: {{.type}}, name: {{.name}}")
	ERR_HOOK_INSTANCE_NOT_INITALED     = errors.TN(SPIRIT_ERR_NS, 39, "hook instance not initaled, hook name: {{.name}}")
	ERR_INPORT_NOT_BIND_HOOK           = errors.TN(SPIRIT_ERR_NS, 40, "in port of {{.inPortName}} not bind hooks")
	ERR_HOOK_BIG_DATA_OPEN_FILE        = errors.TN(SPIRIT_ERR_NS, 41, "open file error,{{.err}}")
	ERR_HOOK_BIG_DATA_UNMARSHAL_CONF   = errors.TN(SPIRIT_ERR_NS, 42, "envJson unmarshal config error,{{.err}}")
	ERR_HOOK_BIG_DATA_CREATE_REDIS     = errors.TN(SPIRIT_ERR_NS, 43, "create redis error,{{.err}}")
	ERR_HOOK_BIG_DATA_REDIS_GET        = errors.TN(SPIRIT_ERR_NS, 44, "redis get data error,{{.err}}")
	ERR_HOOK_BIG_DATA_REDIS_SET        = errors.TN(SPIRIT_ERR_NS, 45, "redis set data error,{{.err}}")
	ERR_JSON_UNMARSHAL                 = errors.TN(SPIRIT_ERR_NS, 46, "json unmarshal data error,{{.err}}")
	ERR_JSON_MARSHAL                   = errors.TN(SPIRIT_ERR_NS, 47, "json marshal data error,{{.err}}")
	ERR_HOOK_NAME_IS_EMPTY             = errors.TN(SPIRIT_ERR_NS, 48, "hook name is empty")
	ERR_HOOK_DRIVER_TYPE_IS_EMPTY      = errors.TN(SPIRIT_ERR_NS, 49, "hook urn type is empty")
	ERR_HEARTBEAT_NOT_EXIST            = errors.TN(SPIRIT_ERR_NS, 50, "heartbeat of {{.name}} not exist")
	ERR_OPTIONS_KEY_NOT_EXIST          = errors.TN(SPIRIT_ERR_NS, 51, "options key of {{.key}} not exist")
	ERR_OPTIONS_VALUE_TYPE_ERROR       = errors.TN(SPIRIT_ERR_NS, 52, "the key of {{.key}}'s value is {{.value}}, but it was not a type of {{.type}}, the real type was {{.realType}}")
	ERR_ASSET_FILE_NOT_CACHED          = errors.TN(SPIRIT_ERR_NS, 53, "asset file of {{.fileName}} not cached")
	ERR_ASSET_ALREADY_LOADED           = errors.TN(SPIRIT_ERR_NS, 54, "asset file already loaded, but did not have same md5: {{.fileName}}")
	ERR_CACHE_FILE_NOT_SPECIFIC        = errors.TN(SPIRIT_ERR_NS, 55, "cache file not specific")
	ERR_CACHE_FILE_LOAD_FAILED         = errors.TN(SPIRIT_ERR_NS, 56, "load cache file failed, file: {{.fileName}}, err: {{.err}}")
	ERR_CACHE_ASSET_BACKUP_FAILED      = errors.TN(SPIRIT_ERR_NS, 57, "backup cache asset failed, file: {{.fileName}}, err: {{.err}}")
	ERR_CACHE_FILE_COULD_NOT_SAVE      = errors.TN(SPIRIT_ERR_NS, 58, "cache file could not save,file: {{.fileName}}, err: {{.err}}")
	ERR_CACHE_FILE_SERIALIZE_FAILED    = errors.TN(SPIRIT_ERR_NS, 59, "cache file serialize failed, file: {{.fileName}}, err: {{.err}}")
	ERR_INSTANCE_HOME_CREATE_FAILED    = errors.TN(SPIRIT_ERR_NS, 60, "could not create instance home: {{.dir}}")
	ERR_SAVE_INSTANCE_METADATA_FAILED  = errors.TN(SPIRIT_ERR_NS, 61, "save instance metadata failed, instance: {{.name}}, err: {{.err}}")
	ERR_INSTANCE_NAME_IS_EMPTY         = errors.TN(SPIRIT_ERR_NS, 62, "instance name is empty")
	ERR_LOAD_INSTANCE_METADATA_FAILED  = errors.TN(SPIRIT_ERR_NS, 63, "load instance metadata failed, you may need recreate the spirit instance, instance: {{.name}}, err: {{.err}}")
	ERR_INSTANCE_CAN_NOT_DEL_WHILE_RUN = errors.TN(SPIRIT_ERR_NS, 64, "instance of {{.name}} cloud not delete while running, running pid: {{.pid}}")
	ERR_OPTIONS_VAL_TYPE_CONV_FAILED   = errors.TN(SPIRIT_ERR_NS, 65, "the key of {{.key}}'s value is {{.value}}, but it was not a type of {{.type}}, and convert it from {{.realType}} failed, err: {{.err}}")
	ERR_INSTANCE_HASH_IS_EMPTY         = errors.TN(SPIRIT_ERR_NS, 66, "the instance of {{.name}} hash is empty, you may recreate this instance")
	ERR_INSTANCE_BIN_ALREADY_COMMIT    = errors.TN(SPIRIT_ERR_NS, 67, "instance of {{.name}} already commited, hash: {{.hash}}")
	ERR_GET_INSTNACE_COMMIT_INFO_ERROR = errors.TN(SPIRIT_ERR_NS, 68, "get instance of {{.name}} commit info error, err: {{.err}}")
	ERR_WRITE_COMMIT_MESSAGE_ERROR     = errors.TN(SPIRIT_ERR_NS, 69, "write commit message error: {{.err}}")
	ERR_MUST_CHECKOUT_BEFORE_REMOVE    = errors.TN(SPIRIT_ERR_NS, 70, "cloud not remove the binary with hash {{.hash}} while using, you can checkout other's bin before remove")
	ERR_REMOVE_BIN_ERROR               = errors.TN(SPIRIT_ERR_NS, 71, "remove binary with hash {{.hash}} failed, err: {{.err}}")
	ERR_COULD_NOT_MAKE_SPIRIT_TMP_DIR  = errors.TN(SPIRIT_ERR_NS, 72, "could not make tmp dir, error: {{.err}}")
	ERR_WRITE_TMP_CONFIG_FILE_FAILED   = errors.TN(SPIRIT_ERR_NS, 73, "write tmp config of {{.fileName}} failed, err: {{.err}}")
	ERR_MARSHAL_CONFIG_FAILED          = errors.TN(SPIRIT_ERR_NS, 74, "marshal config of {{.fileName}} failed, err: {{.err}}")
	ERR_GET_FILE_REL_PATH_FAILED       = errors.TN(SPIRIT_ERR_NS, 75, "get file of {{.fileName}} 's rel path failed, err: {{.err}}")
	ERR_GET_FILE_ABS_PATH_FAILED       = errors.TN(SPIRIT_ERR_NS, 76, "get file of {{.fileName}} 's abs path failed, err: {{.err}}")

	ERR_SPIRIT_NO_CONFIG_SPECIFIC        = errors.TN(SPIRIT_ERR_NS, 1000, "no config file specific")
	ERR_SPIRIT_ENV_FORAMT_ERROR          = errors.TN(SPIRIT_ERR_NS, 1002, "env format error: {{.env}}")
	ERR_SPIRIT_NO_COMPONENTS_TO_RUN      = errors.TN(SPIRIT_ERR_NS, 1003, "no components to run")
	ERR_SPIRIT_COMPONENT_NOT_HOSTING     = errors.TN(SPIRIT_ERR_NS, 1004, "component not hosting: {{.name}}")
	ERR_SPIRIT_COMPONENT_CONF_NOT_EXIST  = errors.TN(SPIRIT_ERR_NS, 1005, "component of {{.name}} did not have configuare in config file")
	ERR_SPIRIT_COMPONENT_CONF_DUP        = errors.TN(SPIRIT_ERR_NS, 1006, "component of {{.name}}'s config duplicated")
	ERR_SPIRIT_SPIRIT_DID_NOT_BUILD      = errors.TN(SPIRIT_ERR_NS, 1007, "spirit did not build")
	ERR_SPIRIT_LOAD_COMPONENT_CONF_ERROR = errors.TN(SPIRIT_ERR_NS, 1008, "load component config error, file: {{.fileName}}, error: {{.err}}")
	ERR_SPIRIT_WRITE_PID_FILE_ERROR      = errors.TN(SPIRIT_ERR_NS, 1009, "write pid file error: {{.err}}")
	ERR_SPIRIT_GET_INSTANCE_PID_FAILED   = errors.TN(SPIRIT_ERR_NS, 1010, "get instance pid of {{.name}} failed, err: {{.err}}")
	ERR_SPIRIT_LIST_INSTANCE_FAILED      = errors.TN(SPIRIT_ERR_NS, 1011, "list instances failed, err: {{.err}}")
	ERR_SPIRIT_READ_PID_FILE_ERROR       = errors.TN(SPIRIT_ERR_NS, 1012, "read instance of {{.name}} pid file error: {{.err}}")
	ERR_SPIRIT_LOCK_PID_FILE_FAILED      = errors.TN(SPIRIT_ERR_NS, 1013, "lock instance pid file failed, instacne: {{.name}}, {{.err}}")
	ERR_SPIRIT_INSTANCE_NAME_NOT_INPUT   = errors.TN(SPIRIT_ERR_NS, 1014, "please input instance name")
	ERR_SPIRIT_ONLY_ONE_INST_NAME        = errors.TN(SPIRIT_ERR_NS, 1015, "only one instance name allowed")
	ERR_SPIRIT_INSTANCE_ALREADY_CREATED  = errors.TN(SPIRIT_ERR_NS, 1016, "the same instance of [{{.name}}] already created, please use start command.")
	ERR_SPIRIT_REMOVE_INSTANCE_FAILED    = errors.TN(SPIRIT_ERR_NS, 1017, "remove spirit instance of [{{.name}}] error: {{.err}}")
	ERR_SPIRIT_INSTANCE_NOT_EXIST        = errors.TN(SPIRIT_ERR_NS, 1018, "spirit instance of {{.name}} not exist")
	ERR_SPIRIT_STOP_INSTANCE_FIALED      = errors.TN(SPIRIT_ERR_NS, 1019, "stop instance of {{.name}} failed, running pid: {{.pid}}, err: {{.err}}")
	ERR_SPIRIT_KILL_INSTANCE_FIALED      = errors.TN(SPIRIT_ERR_NS, 1020, "kill instance of {{.name}} failed, running pid: {{.pid}}, err: {{.err}}")
	ERR_SPIRIT_PAUSE_INSTANCE_FIALED     = errors.TN(SPIRIT_ERR_NS, 1021, "pause or resume instance of {{.name}} failed, running pid: {{.pid}}, err: {{.err}}")
	ERR_SPIRIT_INSTANCE_NOT_RUNNING      = errors.TN(SPIRIT_ERR_NS, 1022, "instance of {{.name}} may not running.")
	ERR_SPIRIT_COPY_INSTANCE_ERROR       = errors.TN(SPIRIT_ERR_NS, 1023, "copy instance error, err: {{.err}}")
	ERR_SPIRIT_START_CHILD_INSTANCE      = errors.TN(SPIRIT_ERR_NS, 1024, "start child instance of {{.name}} failed, err: {{.err}}")
	ERR_SPIRIT_COMMIT_MSG_NOT_INPUT      = errors.TN(SPIRIT_ERR_NS, 1025, "commit message not input")
	ERR_SPIRIT_INSTANCE_HASH_NOT_INPUT   = errors.TN(SPIRIT_ERR_NS, 1026, "spirit instance hash not input")
)
