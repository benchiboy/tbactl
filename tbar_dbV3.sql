#�˻�������Ϣ��
CREATE TABLE `tba_account_comments` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `comment_no` bigint(20) DEFAULT '0' COMMENT '���۱��',
  `user_id` int(11) DEFAULT '0' COMMENT '�������۵��˻�ID',
  `flow_no` int(11) DEFAULT '0' COMMENT '���۵�����',
  `old_comment_no` int(11) DEFAULT  '0'  COMMENT  '�����Ĭ��ֵ�����������Ƹ��λ��������ĳ�����۵Ļظ�',
  `comment_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `comment_desc`  varchar(1000)  default ''   COMMENT '��������',
  `insert_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `update_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `version` int(11) unsigned DEFAULT '0' COMMENT '�汾',
  PRIMARY KEY (`id`),
  KEY `idx_comment_no` (`comment_no`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='�û�������Ϣ��';

#�˻�������ˮ��Ϣ��
CREATE TABLE `tba_account_flows` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned DEFAULT '0' COMMENT '�˻�ID',
  `trxn_no` bigint(20)  NOT NULL  COMMENT '������ˮ��',
  `trxn_date` datetime DEFAULT NULL COMMENT '����ʱ��',
  `trxn_amt` decimal(15,2) DEFAULT '0.00' COMMENT '���׽��',
  `trxn_type` varchar(10) DEFAULT '' COMMENT '�������ͣ������ʽ��ף�������Ʒ����',
  `proc_status` varchar(10) DEFAULT '' COMMENT '���״���״̬ ',
  `proc_msg` varchar(10) DEFAULT '' COMMENT '���״�����ԭ��',
  `goods_no` varchar(10) DEFAULT '' COMMENT '��Ʒ���',
  `discount_rate` decimal(5,3) DEFAULT '0.000' COMMENT '��Ʒ�ۿ۱���',
  `promotion_no`  varchar(10) DEFAULT ''  COMMENT '��������',
  `account_bal` decimal(15,2) DEFAULT '0.00' COMMENT '�˻����',
  `trxn_memo` varchar(50) DEFAULT '' COMMENT '���ױ�ע',
  `done_date` datetime DEFAULT NULL COMMENT '����ȷ��ʱ��',
  `insert_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `update_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `update_user` varchar(50) DEFAULT '' COMMENT '�˹��������׵��û�ID',
  `version` int(11) unsigned DEFAULT '0' COMMENT '�汾',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_trxn_date` (`trxn_date`),
  KEY `idx_proc_status` (`proc_status`),
  KEY `idx_trxn_no` (`trxn_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='�û�������ʷ��';

#�˻���¼��־��ˮ
CREATE TABLE `tba_account_logins` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned DEFAULT '0' COMMENT '�˻�ID',
  `login_no` bigint(20)  NOT NULL  COMMENT '��¼��ˮ��',
  `login_time` datetime DEFAULT NULL COMMENT '��¼ʱ��',
  `login_desc` varchar(30) DEFAULT '' COMMENT '��¼����',
  `login_result` tinyint(4) DEFAULT '0' COMMENT '��¼���',
  `device_ip` varchar(30) DEFAULT '' COMMENT '�豸ip',
  `device_type` tinyint(4) DEFAULT '0' COMMENT '�豸���ͣ�1��ANDROID, 2��OS, 3��PC',
  `device_os` varchar(30) DEFAULT '' COMMENT '�豸����ϵͳ',
  `device_os_ver` varchar(30) DEFAULT '' COMMENT '�豸����ϵͳ�汾',
  `device_id` varchar(30) DEFAULT '' COMMENT '�豸id',
  `latitude` varchar(20) DEFAULT '' COMMENT 'γ��',
  `longitude` varchar(20) DEFAULT '' COMMENT '����',
  `insert_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `version` int(11) unsigned DEFAULT '0' COMMENT '�汾',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_login_no` (`login_no`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='�û���¼�˺ű�';

#�˻�������־��ˮ
CREATE TABLE `tba_account_actions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned DEFAULT '0' COMMENT '�˻�ID',
  `login_no` bigint(20)  NOT NULL  COMMENT '��¼��ˮ��',
  `action_no` bigint(20)  unsigned DEFAULT '0'  COMMENT '�û�������ˮ��',
  `action_type`  tinyint(4) DEFAULT '0'  COMMENT '�������ͣ�1���ֶ��ϣ�2:��ť��',
  `flow_no`  int(11) unsigned DEFAULT '0' COMMENT '���̱��',
  `service_point_no` smallint(6)    DEFAULT '0' COMMENT '�������',
  `field_name`  varchar(30) DEFAULT '20'  COMMENT '�ֶ�����',
  `duration`   smallint(6)    DEFAULT '0'    COMMENT 'ͣ��ʱ�䣬��λ��',
  `insert_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `version` int(11) unsigned DEFAULT '0' COMMENT '�汾',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_login_no` (`login_no`),
  KEY `idx_action_no` (`action_no`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='�˻�������־��ˮ';

#�˻��ҵĲ�����¼��
CREATE TABLE `tba_account_posts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_no` bigint(20) DEFAULT '0' COMMENT 'Ͷ�ͼ�¼���',
  `user_id` int(11) unsigned DEFAULT '0' COMMENT '�������˻�ID',
  `src_user_id` int(11) DEFAULT '0'  COMMENT '�����Ķ����˻�ID',
  `dst_user_id` int(11) DEFAULT '0'  COMMENT 'Ŀ���˻�ID',
  `flow_no`  int(11) DEFAULT '0' COMMENT '���̱�ű��',
  `insert_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `update_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `version` int(11) unsigned DEFAULT '0' COMMENT '�汾',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`src_user_id`),
  KEY `idx_post_no (`post_no`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='�˻��ҵĲ�����¼��';

#�˻���Ϣ��
DROP TABLE tba_accounts;
CREATE TABLE `tba_accounts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned DEFAULT '0' COMMENT '��¼�˻�ID',
  `partner_user_id` varchar(36)  DEFAULT '' COMMENT '���������û�ID',
  `parent_user_id` int(11) DEFAULT '0' COMMENT '�Ƽ��û�ID',
  `user_role` tinyint(4) DEFAULT '0' COMMENT '�û���ǰ��ɫ 1����ְ�� 2��ӦƸ��  3���ϻ���',
  `user_status` tinyint(4) DEFAULT '0' COMMENT '��ǰ�û�״̬: 0�� ���� 1�������������,2���˻��˹����ᣨ�ʽ𲻿ɽ�����,3 :�˻�ֹ�����˻����������ѣ�, 4���˻�ֹ �루�˻����ʳ�ֵ��',
  `avatar_url` varchar(100) DEFAULT '' COMMENT '�û�ͷ��URL',
  `login_mode` tinyint(4) DEFAULT '0' COMMENT '��¼��ʽ��1��΢�ţ�2���ֻ�ע�ᣬ3��΢��',
  `login_name` varchar(50) DEFAULT '' COMMENT '��¼�˺�',
  `login_password` varchar(50) DEFAULT '' COMMENT '����',
  `error_count` tinyint(4) DEFAULT '0' COMMENT '����������',
  `last_login_time` datetime DEFAULT NULL COMMENT '�ϴε�¼ʱ��',
  `last_device_id`    varchar(30) DEFAULT ''  COMMENT '�ϴε�¼�豸ID',
  `account_bal` decimal(15,2) DEFAULT '0.00' COMMENT '�˻����-�ʽ�',
  `goods_count` decimal(15,2) DEFAULT '0.00' COMMENT '��Ʒ����-����',
  `device_ip` varchar(30) DEFAULT '' COMMENT ' �״��豸ip',
  `device_type` tinyint(4) DEFAULT '0' COMMENT '�״��豸���ͣ�1��ANDROID, 2��OS, 3��PC',
  `device_os` varchar(30) DEFAULT '' COMMENT '�״��豸����ϵͳ',
  `device_os_ver` varchar(30) DEFAULT '' COMMENT '�״��豸����ϵͳ�汾',
  `device_id` varchar(30) DEFAULT '' COMMENT '�״��豸id',
  `latitude` varchar(20) DEFAULT '' COMMENT '�״�γ��',
  `longitude` varchar(20) DEFAULT '' COMMENT '�״ξ���',
  `market` varchar(30) DEFAULT '' COMMENT 'Ӧ���г�',
  `user_channel` varchar(20) DEFAULT '' COMMENT '�������',
  `random_no` int(11) DEFAULT '0' COMMENT '�û������',
  `region_no`   varchar(15) DEFAULT '' COMMENT '�û������������',
  `customer_id` int(11) unsigned DEFAULT '0' COMMENT '�ͻ�ID',
  `created_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `updated_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `memo` varchar(50) DEFAULT '' COMMENT '��ע�ֶ�',
  `insert_user` varchar(50) DEFAULT '' COMMENT '�����û�',
  `update_user` varchar(50) DEFAULT '' COMMENT '�����û�',
  `version` int(11) unsigned DEFAULT '0' COMMENT '�汾',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_partner_user_id` (`partner_user_id`),
  KEY `idx_user_status` (`user_status`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='�˻���Ϣ��';


#��ְ����Ϣ��
CREATE TABLE `tba_customers` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) unsigned DEFAULT '0' COMMENT '��ְ��ID',
  `customer_type`  tinyint(4) DEFAULT '0' COMMENT '�û����ͣ�1����ְ�� 2��ӦƸ��  3���ϻ���',
  `customer_no` varchar(21) DEFAULT '0' COMMENT '֤������',
  `customer_name` varchar(200) DEFAULT '0' COMMENT '�û���ʵ����',
  `sex`  tinyint(4)  DEFAULT '0' COMMENT '1:�У�2��Ů',
  `mail` varchar(20) DEFAULT '' COMMENT '����',
  `url` varchar(50) DEFAULT '' COMMENT '���˻���ҵ��ҳ',
  `customer_desc` varchar(500) DEFAULT '' COMMENT '���˻����',
  `update_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `version`  int(11)  unsigned DEFAULT '0' COMMENT '�汾',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`customer_id`)
  
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='��ְ����Ϣ��';


#��ְ�ߣ����ˣ�Ͷ�ͼ�����¼��
drop table tba_account_post_resumes;
CREATE TABLE `tba_account_post_resumes` (
  `id`   bigint(20) NOT NULL AUTO_INCREMENT,
  `post_no`    bigint(20)  DEFAULT '0' COMMENT 'Ͷ�ͼ�¼���',
  `user_id`   int(11)  unsigned  DEFAULT '0' 	COMMENT 'Ͷ�ͼ������û�ID',
  `user_name` varchar(50)  DEFAULT '' COMMENT '��ְ������',
  `user_sex`   tinyint(4) DEFAULT '0'  COMMENT '��ְ���Ա�:1 �� 2��Ů',
  `user_phone` varchar(21)  DEFAULT '' COMMENT '�ֻ�����',
  `work_years`  varchar(10)  DEFAULT '' COMMENT '��������',
  `edu_level`  varchar(10)  DEFAULT '' COMMENT '���ѧ��',
  `want_position_no` varchar(10)  DEFAULT '' COMMENT '��ְ������ְλ',
  `want_salary` varchar(10)   DEFAULT ''  COMMENT '��ְ������н��',
  `want_area`  varchar(10)  DEFAULT ''  COMMENT '��ְ����������',
  `insert_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `update_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `version` int(11)  unsigned DEFAULT '0' COMMENT '�汾',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_post_no` (`post_no`)
  ) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='��ְ�ߣ����ˣ�Ͷ�ͼ�����¼��';

#��ְ�߽������
CREATE TABLE `tba_account_educations` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id`   int(11)  unsigned DEFAULT '0' COMMENT '�û�ID',
  `post_no`    bigint(20)  DEFAULT '0' COMMENT 'Ͷ�ͼ�¼���',
  `school_name` varchar(50) DEFAULT '' COMMENT 'ѧУ����',
  `school_level`   tinyint(4)   DEFAULT '0' COMMENT '��ҵ���� 1����ʿ��2���о�����3�����ơ�4����ר��5����ר��6�����С�7��Сѧ��9������',
  `school_major` varchar(10) DEFAULT '' COMMENT '��ѧרҵ��רҵ�����',
  `school_end` date DEFAULT NULL COMMENT '��ҵʱ��',
  `insert_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `update_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `version` int(11)  unsigned DEFAULT '0' COMMENT '�汾',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_post_no` (`post_no`),
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='��ְ�߽������'

#��ְ����Ŀ�����
CREATE TABLE `tba_account_projects` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11)  unsigned DEFAULT '0' COMMENT '�û�ID',
  `post_no`    bigint(20)  DEFAULT '0' COMMENT 'Ͷ�ͼ�¼���',
  `project_name` varchar(100) DEFAULT '' COMMENT '��Ŀ����',
  `project_begin` date DEFAULT NULL COMMENT '�μ���Ŀ��ʼ����',
  `project_end` date DEFAULT NULL COMMENT '�μ���Ŀ��������',
  `project_desc` varchar(500) DEFAULT '' COMMENT '�Լ���������',
  `project_url1` varchar(200) DEFAULT '' COMMENT '��Ŀ������ͼƬ����',
  `project_url2` varchar(200) DEFAULT '' COMMENT '��Ŀ�ɹ�URL',
  `project_url3` varchar(200) DEFAULT '' COMMENT '��Ŀ�ɹ�URL',
   `insert_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `update_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `version` int(11)  unsigned DEFAULT '0' COMMENT '�汾',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_post_no` (`post_no`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='��ְ����Ŀ�����';


#��ְ��������������
CREATE TABLE `tba_account_others` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11)  unsigned DEFAULT '0' COMMENT '�û�ID',
  `post_no`    bigint(20)  DEFAULT '0' COMMENT 'Ͷ�ͼ�¼���',
  `other_desc` varchar(500) DEFAULT '' COMMENT '�Լ���������',
  `other_url1` varchar(200) DEFAULT '' COMMENT '��������������ͼƬ1',
  `other_url2` varchar(200) DEFAULT '' COMMENT '��������������ͼƬ2',
  `other_url3` varchar(200) DEFAULT '' COMMENT '��������������ͼƬ3',
  `insert_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `update_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `version` int(11)  unsigned DEFAULT '0' COMMENT '�汾',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_post_no` (`post_no`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='��ְ��������������';


#��Ƹ�߷���ְλ��¼
CREATE TABLE `tba_account_post_positions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `publish_no`  bigint(20)  DEFAULT '0' COMMENT 'ְλ�������',
  `user_id`  	int(11)  DEFAULT '0' COMMENT '�û�ID',
  `position_name` varchar(50) DEFAULT '' COMMENT 'ְλ����',
  `position_desc` varchar(1000)  DEFAULT '' COMMENT 'ְλ��ְҪ��',
  `publish_time` datetime DEFAULT null  COMMENT 'ְλ����ʱ��',
  `expire_time` date  DEFAULT NULL COMMENT 'ְλ������',
  `salary_min`  decimal(15,2) DEFAULT '0' COMMENT '��Снˮ',
  `salary_max`  decimal(15,2) DEFAULT '0' COMMENT '���нˮ',
  `city` varchar(50) DEFAULT '' COMMENT 	'ְλҪ�����',
  `school_level` varchar(10) DEFAULT '' COMMENT 'Ҫ��ѧ��',
  `exp_min` varchar(10)  DEFAULT NULL COMMENT 'Ҫ����С��������',
  `exp_max` varchar(10)   DEFAULT '' COMMENT 'Ҫ�����������',
  `insert_time` datetime DEFAULT  NULL COMMENT '����ʱ��',
  `update_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `version` int(11)  unsigned DEFAULT '0' COMMENT '�汾',
  PRIMARY KEY (`id`),
  KEY `idx_publish_no` (`publish_no`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='��Ƹ�߷���ְλ��¼';

#�ϻ����Ƽ���ʷ��¼��
CREATE TABLE `tba_account_recommends` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11)  unsigned   DEFAULT '0' COMMENT '������ķ�����',
  `recommend_no`  bigint(20)  DEFAULT '0' COMMENT 'ְλ�������',
  `recommend_time` datetime DEFAULT NULL COMMENT '�ƹ�ʱ��',
  `recommend_type` tinyint(4) DEFAULT '0' COMMENT '1:ע���Ƽ���2�����͸�λ����ְ�ߣ�3��������ְ������Ƹ��',
  `insert_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `update_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `version` int(11)  unsigned DEFAULT '0' COMMENT '�汾',
   PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='�ϻ���������ʷ��';