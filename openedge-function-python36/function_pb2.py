# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: function.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='function.proto',
  package='openedge',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0e\x66unction.proto\x12\x08openedge\"z\n\x0f\x46unctionMessage\x12\n\n\x02ID\x18\x01 \x01(\x04\x12\x0b\n\x03QOS\x18\x02 \x01(\r\x12\r\n\x05Topic\x18\x03 \x01(\t\x12\x0f\n\x07Payload\x18\x04 \x01(\x0c\x12\x14\n\x0c\x46unctionName\x18\x0b \x01(\t\x12\x18\n\x10\x46unctionInvokeID\x18\x0c \x01(\t2J\n\x08\x46unction\x12>\n\x04\x43\x61ll\x12\x19.openedge.FunctionMessage\x1a\x19.openedge.FunctionMessage\"\x00\x62\x06proto3')
)




_FUNCTIONMESSAGE = _descriptor.Descriptor(
  name='FunctionMessage',
  full_name='openedge.FunctionMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ID', full_name='openedge.FunctionMessage.ID', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='QOS', full_name='openedge.FunctionMessage.QOS', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Topic', full_name='openedge.FunctionMessage.Topic', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Payload', full_name='openedge.FunctionMessage.Payload', index=3,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='FunctionName', full_name='openedge.FunctionMessage.FunctionName', index=4,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='FunctionInvokeID', full_name='openedge.FunctionMessage.FunctionInvokeID', index=5,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=28,
  serialized_end=150,
)

DESCRIPTOR.message_types_by_name['FunctionMessage'] = _FUNCTIONMESSAGE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

FunctionMessage = _reflection.GeneratedProtocolMessageType('FunctionMessage', (_message.Message,), dict(
  DESCRIPTOR = _FUNCTIONMESSAGE,
  __module__ = 'function_pb2'
  # @@protoc_insertion_point(class_scope:openedge.FunctionMessage)
  ))
_sym_db.RegisterMessage(FunctionMessage)



_FUNCTION = _descriptor.ServiceDescriptor(
  name='Function',
  full_name='openedge.Function',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=152,
  serialized_end=226,
  methods=[
  _descriptor.MethodDescriptor(
    name='Call',
    full_name='openedge.Function.Call',
    index=0,
    containing_service=None,
    input_type=_FUNCTIONMESSAGE,
    output_type=_FUNCTIONMESSAGE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_FUNCTION)

DESCRIPTOR.services_by_name['Function'] = _FUNCTION

# @@protoc_insertion_point(module_scope)
