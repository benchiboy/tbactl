#账户评论信息表
CREATE TABLE `tba_account_comments` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `comment_no` bigint(20) DEFAULT '0' COMMENT '评论编号',
  `user_id` int(11) DEFAULT '0' COMMENT '发表评论的账户ID',
  `flow_no` int(11) DEFAULT '0' COMMENT '评论的流程',
  `old_comment_no` int(11) DEFAULT  '0'  COMMENT  '如果是默认值，评论针对招聘岗位，否则是某个评论的回复',
  `comment_time` datetime DEFAULT NULL COMMENT '评论时间',
  `comment_desc`  varchar(1000)  default ''   COMMENT '评论描述',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11) unsigned DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`id`),
  KEY `idx_comment_no` (`comment_no`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='用户评论信息表';

#账户交易流水信息表
CREATE TABLE `tba_account_flows` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned DEFAULT '0' COMMENT '账户ID',
  `trxn_no` bigint(20)  NOT NULL  COMMENT '交易流水号',
  `trxn_date` datetime DEFAULT NULL COMMENT '交易时间',
  `trxn_amt` decimal(15,2) DEFAULT '0.00' COMMENT '交易金额',
  `trxn_type` varchar(10) DEFAULT '' COMMENT '交易类型，包括资金交易，虚拟商品交易',
  `proc_status` varchar(10) DEFAULT '' COMMENT '交易处理状态 ',
  `proc_msg` varchar(10) DEFAULT '' COMMENT '交易处理结果原因',
  `goods_no` varchar(10) DEFAULT '' COMMENT '商品编号',
  `discount_rate` decimal(5,3) DEFAULT '0.000' COMMENT '商品折扣比例',
  `promotion_no`  varchar(10) DEFAULT ''  COMMENT '促销活动编号',
  `account_bal` decimal(15,2) DEFAULT '0.00' COMMENT '账户余额',
  `trxn_memo` varchar(50) DEFAULT '' COMMENT '交易备注',
  `done_date` datetime DEFAULT NULL COMMENT '交易确认时间',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `update_user` varchar(50) DEFAULT '' COMMENT '人工调整交易的用户ID',
  `version` int(11) unsigned DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_trxn_date` (`trxn_date`),
  KEY `idx_proc_status` (`proc_status`),
  KEY `idx_trxn_no` (`trxn_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户交易历史表';

#账户登录日志流水
CREATE TABLE `tba_account_logins` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned DEFAULT '0' COMMENT '账户ID',
  `login_no` bigint(20)  NOT NULL  COMMENT '登录流水号',
  `login_time` datetime DEFAULT NULL COMMENT '登录时间',
  `login_desc` varchar(30) DEFAULT '' COMMENT '登录描述',
  `login_result` tinyint(4) DEFAULT '0' COMMENT '登录结果',
  `device_ip` varchar(30) DEFAULT '' COMMENT '设备ip',
  `device_type` tinyint(4) DEFAULT '0' COMMENT '设备类型：1：ANDROID, 2：OS, 3：PC',
  `device_os` varchar(30) DEFAULT '' COMMENT '设备操作系统',
  `device_os_ver` varchar(30) DEFAULT '' COMMENT '设备操作系统版本',
  `device_id` varchar(30) DEFAULT '' COMMENT '设备id',
  `latitude` varchar(20) DEFAULT '' COMMENT '纬度',
  `longitude` varchar(20) DEFAULT '' COMMENT '经度',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `version` int(11) unsigned DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_login_no` (`login_no`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='用户登录账号表';

#账户操作日志流水
CREATE TABLE `tba_account_actions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned DEFAULT '0' COMMENT '账户ID',
  `login_no` bigint(20)  NOT NULL  COMMENT '登录流水号',
  `action_no` bigint(20)  unsigned DEFAULT '0'  COMMENT '用户操作流水号',
  `action_type`  tinyint(4) DEFAULT '0'  COMMENT '动作类型：1：字段上，2:按钮上',
  `flow_no`  int(11) unsigned DEFAULT '0' COMMENT '流程编号',
  `service_point_no` smallint(6)    DEFAULT '0' COMMENT '服务点编号',
  `field_name`  varchar(30) DEFAULT '20'  COMMENT '字段名称',
  `duration`   smallint(6)    DEFAULT '0'    COMMENT '停留时间，单位秒',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `version` int(11) unsigned DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_login_no` (`login_no`),
  KEY `idx_action_no` (`action_no`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='账户操作日志流水';

#账户我的操作记录表
CREATE TABLE `tba_account_posts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_no` bigint(20) DEFAULT '0' COMMENT '投送记录编号',
  `user_id` int(11) unsigned DEFAULT '0' COMMENT '操作的账户ID',
  `src_user_id` int(11) DEFAULT '0'  COMMENT '操作的对象账户ID',
  `dst_user_id` int(11) DEFAULT '0'  COMMENT '目的账户ID',
  `flow_no`  int(11) DEFAULT '0' COMMENT '流程编号编号',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11) unsigned DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`src_user_id`),
  KEY `idx_post_no (`post_no`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='账户我的操作记录表';

#账户信息表
DROP TABLE tba_accounts;
CREATE TABLE `tba_accounts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned DEFAULT '0' COMMENT '登录账户ID',
  `partner_user_id` varchar(36)  DEFAULT '' COMMENT '第三方的用户ID',
  `parent_user_id` int(11) DEFAULT '0' COMMENT '推荐用户ID',
  `user_role` tinyint(4) DEFAULT '0' COMMENT '用户当前角色 1：求职者 2：应聘者  3：合伙人',
  `user_status` tinyint(4) DEFAULT '0' COMMENT '当前用户状态: 0： 正常 1：密码错误锁定,2：账户人工冻结（资金不可进出）,3 :账户止付（账户不允许消费）, 4：账户止 入（账户不允充值）',
  `avatar_url` varchar(100) DEFAULT '' COMMENT '用户头像URL',
  `login_mode` tinyint(4) DEFAULT '0' COMMENT '登录方式：1：微信，2：手机注册，3：微博',
  `login_name` varchar(50) DEFAULT '' COMMENT '登录账号',
  `login_password` varchar(50) DEFAULT '' COMMENT '密码',
  `error_count` tinyint(4) DEFAULT '0' COMMENT '密码错误次数',
  `last_login_time` datetime DEFAULT NULL COMMENT '上次登录时间',
  `last_device_id`    varchar(30) DEFAULT ''  COMMENT '上次登录设备ID',
  `account_bal` decimal(15,2) DEFAULT '0.00' COMMENT '账户余额-资金',
  `goods_count` decimal(15,2) DEFAULT '0.00' COMMENT '商品数量-虚拟',
  `device_ip` varchar(30) DEFAULT '' COMMENT ' 首次设备ip',
  `device_type` tinyint(4) DEFAULT '0' COMMENT '首次设备类型：1：ANDROID, 2：OS, 3：PC',
  `device_os` varchar(30) DEFAULT '' COMMENT '首次设备操作系统',
  `device_os_ver` varchar(30) DEFAULT '' COMMENT '首次设备操作系统版本',
  `device_id` varchar(30) DEFAULT '' COMMENT '首次设备id',
  `latitude` varchar(20) DEFAULT '' COMMENT '首次纬度',
  `longitude` varchar(20) DEFAULT '' COMMENT '首次经度',
  `market` varchar(30) DEFAULT '' COMMENT '应用市场',
  `user_channel` varchar(20) DEFAULT '' COMMENT '获客渠道',
  `random_no` int(11) DEFAULT '0' COMMENT '用户随机数',
  `region_no`   varchar(15) DEFAULT '' COMMENT '用户负责的区域编号',
  `customer_id` int(11) unsigned DEFAULT '0' COMMENT '客户ID',
  `created_time` datetime DEFAULT NULL COMMENT '插入时间',
  `updated_time` datetime DEFAULT NULL COMMENT '更新时间',
  `memo` varchar(50) DEFAULT '' COMMENT '备注字段',
  `insert_user` varchar(50) DEFAULT '' COMMENT '插入用户',
  `update_user` varchar(50) DEFAULT '' COMMENT '更新用户',
  `version` int(11) unsigned DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_partner_user_id` (`partner_user_id`),
  KEY `idx_user_status` (`user_status`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='账户信息表';


#求职者信息表
CREATE TABLE `tba_customers` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) unsigned DEFAULT '0' COMMENT '求职者ID',
  `customer_type`  tinyint(4) DEFAULT '0' COMMENT '用户类型：1：求职者 2：应聘者  3：合伙人',
  `customer_no` varchar(21) DEFAULT '0' COMMENT '证件号码',
  `customer_name` varchar(200) DEFAULT '0' COMMENT '用户真实姓名',
  `sex`  tinyint(4)  DEFAULT '0' COMMENT '1:男，2：女',
  `mail` varchar(20) DEFAULT '' COMMENT '邮箱',
  `url` varchar(50) DEFAULT '' COMMENT '个人或企业主页',
  `customer_desc` varchar(500) DEFAULT '' COMMENT '个人或介绍',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version`  int(11)  unsigned DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`customer_id`)
  
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='求职者信息表';


#求职者（个人）投送简历记录表
drop table tba_account_post_resumes;
CREATE TABLE `tba_account_post_resumes` (
  `id`   bigint(20) NOT NULL AUTO_INCREMENT,
  `post_no`    bigint(20)  DEFAULT '0' COMMENT '投送记录编号',
  `user_id`   int(11)  unsigned  DEFAULT '0' 	COMMENT '投送简历的用户ID',
  `user_name` varchar(50)  DEFAULT '' COMMENT '求职者姓名',
  `user_sex`   tinyint(4) DEFAULT '0'  COMMENT '求职者性别:1 男 2：女',
  `user_phone` varchar(21)  DEFAULT '' COMMENT '手机号码',
  `work_years`  varchar(10)  DEFAULT '' COMMENT '工作年限',
  `edu_level`  varchar(10)  DEFAULT '' COMMENT '最高学历',
  `want_position_no` varchar(10)  DEFAULT '' COMMENT '求职期望的职位',
  `want_salary` varchar(10)   DEFAULT ''  COMMENT '求职期望的薪资',
  `want_area`  varchar(10)  DEFAULT ''  COMMENT '求职期望的区域',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11)  unsigned DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_post_no` (`post_no`)
  ) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='求职者（个人）投送简历记录表';

#求职者教育情况
CREATE TABLE `tba_account_educations` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id`   int(11)  unsigned DEFAULT '0' COMMENT '用户ID',
  `post_no`    bigint(20)  DEFAULT '0' COMMENT '投送记录编号',
  `school_name` varchar(50) DEFAULT '' COMMENT '学校名称',
  `school_level`   tinyint(4)   DEFAULT '0' COMMENT '毕业级别： 1：博士、2：研究生、3：本科、4：大专、5：中专、6：高中、7：小学、9：其他',
  `school_major` varchar(10) DEFAULT '' COMMENT '所学专业：专业代码表',
  `school_end` date DEFAULT NULL COMMENT '毕业时间',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11)  unsigned DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_post_no` (`post_no`),
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='求职者教育情况'

#求职者项目经验表
CREATE TABLE `tba_account_projects` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11)  unsigned DEFAULT '0' COMMENT '用户ID',
  `post_no`    bigint(20)  DEFAULT '0' COMMENT '投送记录编号',
  `project_name` varchar(100) DEFAULT '' COMMENT '项目名称',
  `project_begin` date DEFAULT NULL COMMENT '参加项目开始日期',
  `project_end` date DEFAULT NULL COMMENT '参加项目结束日期',
  `project_desc` varchar(500) DEFAULT '' COMMENT '自己工作描述',
  `project_url1` varchar(200) DEFAULT '' COMMENT '项目语音或图片介绍',
  `project_url2` varchar(200) DEFAULT '' COMMENT '项目成果URL',
  `project_url3` varchar(200) DEFAULT '' COMMENT '项目成果URL',
   `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11)  unsigned DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_post_no` (`post_no`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='求职者项目经验表';


#求职者其他附件材料
CREATE TABLE `tba_account_others` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11)  unsigned DEFAULT '0' COMMENT '用户ID',
  `post_no`    bigint(20)  DEFAULT '0' COMMENT '投送记录编号',
  `other_desc` varchar(500) DEFAULT '' COMMENT '自己工作描述',
  `other_url1` varchar(200) DEFAULT '' COMMENT '个人其他附件，图片1',
  `other_url2` varchar(200) DEFAULT '' COMMENT '个人其他附件，图片2',
  `other_url3` varchar(200) DEFAULT '' COMMENT '个人其他附件，图片3',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11)  unsigned DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_post_no` (`post_no`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='求职者其他附件材料';


#招聘者发布职位记录
CREATE TABLE `tba_account_post_positions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `publish_no`  bigint(20)  DEFAULT '0' COMMENT '职位发布编号',
  `user_id`  	int(11)  DEFAULT '0' COMMENT '用户ID',
  `position_name` varchar(50) DEFAULT '' COMMENT '职位名称',
  `position_desc` varchar(1000)  DEFAULT '' COMMENT '职位任职要求',
  `publish_time` datetime DEFAULT null  COMMENT '职位发布时间',
  `expire_time` date  DEFAULT NULL COMMENT '职位到期日',
  `salary_min`  decimal(15,2) DEFAULT '0' COMMENT '最小薪水',
  `salary_max`  decimal(15,2) DEFAULT '0' COMMENT '最大薪水',
  `city` varchar(50) DEFAULT '' COMMENT 	'职位要求城市',
  `school_level` varchar(10) DEFAULT '' COMMENT '要求学历',
  `exp_min` varchar(10)  DEFAULT NULL COMMENT '要求最小工作年限',
  `exp_max` varchar(10)   DEFAULT '' COMMENT '要求最大工作年限',
  `insert_time` datetime DEFAULT  NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11)  unsigned DEFAULT '0' COMMENT '版本',
  PRIMARY KEY (`id`),
  KEY `idx_publish_no` (`publish_no`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='招聘者发布职位记录';

#合伙者推荐历史记录表
CREATE TABLE `tba_account_recommends` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11)  unsigned   DEFAULT '0' COMMENT '推送码的发起者',
  `recommend_no`  bigint(20)  DEFAULT '0' COMMENT '职位发布编号',
  `recommend_time` datetime DEFAULT NULL COMMENT '推广时间',
  `recommend_type` tinyint(4) DEFAULT '0' COMMENT '1:注册推荐，2：推送岗位至求职者，3：推送求职者至招聘者',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `version` int(11)  unsigned DEFAULT '0' COMMENT '版本',
   PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='合伙者推送历史表';