Resources:
  default:
    Type: TencentCloud::Serverless::Namespace
    scfgw:   # 函数名改这里 需要到控制台添加函数
      Type: TencentCloud::Serverless::Function
      Properties:
        CodeUri: ./
        Type: Event
        Description: This is a template function
        # Role: QCS_SCFExcuteRole
        Environment:
          Variables:
            ENV_FIRST: env1
            ENV_SECOND: env2
        Handler: index
        MemorySize: 128
        Runtime: Go1
        Timeout: 3
        #VpcConfig:
        #    VpcId: 'vpc-qdqc5k2p'
        #    SubnetId: 'subnet-pad6l61i'
        #Events:
        #    timer:
        #        Type: Timer
        #        Properties:
        #            CronExpression: '*/5 * * * *'
        #            Enable: True
        #    cli-appid.cos.ap-beijing.myqcloud.com: # full bucket name
        #        Type: COS
        #        Properties:
        #            Bucket: cli-appid.cos.ap-beijing.myqcloud.com
        #            Filter:
        #                Prefix: filterdir/
        #                Suffix: .jpg
        #            Events: cos:ObjectCreated:*
        #            Enable: True
        #    topic:            # topic name
        #        Type: CMQ
        #        Properties:
        #            Name: qname
        #    hello_world_apigw:  # ${FunctionName} + '_apigw'
        #        Type: APIGW
        #        Properties:
        #            StageName: release
        #            ServiceId:
        #            HttpMethod: ANY

Globals:
  Function:
    Timeout: 10
