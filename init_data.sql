-- 删除所有表
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS `groups`;
DROP TABLE IF EXISTS group_member;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS tasks_record;
DROP TABLE IF EXISTS sign_in_records;
DROP TABLE IF EXISTS check_application;
DROP TABLE IF EXISTS join_application;

-- 创建 users 表
CREATE TABLE users (
    user_id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(128) NOT NULL,
    mail VARCHAR(128) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id),
    UNIQUE INDEX idx_username (username),
    UNIQUE INDEX idx_mail (mail)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 清空所有表数据
SET FOREIGN_KEY_CHECKS = 0;
TRUNCATE TABLE users;
TRUNCATE TABLE `groups`;
TRUNCATE TABLE group_member;
TRUNCATE TABLE tasks;
TRUNCATE TABLE tasks_record;
TRUNCATE TABLE sign_in_records;
TRUNCATE TABLE check_application;
TRUNCATE TABLE join_application;
SET FOREIGN_KEY_CHECKS = 1;

-- 插入用户数据 (密码: 12345678 的MD5值)
INSERT INTO users (username, password, mail) VALUES
('游辰昊', '25d55ad283aa400af464c76d713c07ad', 'youchenhao@example.com'),
('曹航语', '25d55ad283aa400af464c76d713c07ad', 'caohangyu@example.com'),
('刘胜杰', '25d55ad283aa400af464c76d713c07ad', 'liushengjie@example.com'),
('张明', '25d55ad283aa400af464c76d713c07ad', 'zhangming@example.com'),
('李华', '25d55ad283aa400af464c76d713c07ad', 'lihua@example.com'),
('王芳', '25d55ad283aa400af464c76d713c07ad', 'wangfang@example.com');

-- 插入小组数据
INSERT INTO `groups` (group_name, description, creator_id, creator_name, member_num) VALUES
('编译原理课程组', '编译原理课程学习小组', 1, '游辰昊', 3),
('实训项目组', '企业实训项目开发小组', 2, '曹航语', 4),
('外勤工作组', '公司外勤工作小组', 3, '刘胜杰', 3);

-- 插入小组成员数据
INSERT INTO group_member (group_id, user_id, group_name, username, role) VALUES
-- 编译原理课程组
(1, 1, '编译原理课程组', '游辰昊', 'admin'),
(1, 2, '编译原理课程组', '曹航语', 'member'),
(1, 3, '编译原理课程组', '刘胜杰', 'member'),
-- 实训项目组
(2, 2, '实训项目组', '曹航语', 'admin'),
(2, 1, '实训项目组', '游辰昊', 'member'),
(2, 4, '实训项目组', '张明', 'member'),
(2, 5, '实训项目组', '李华', 'member'),
-- 外勤工作组
(3, 3, '外勤工作组', '刘胜杰', 'admin'),
(3, 4, '外勤工作组', '张明', 'member'),
(3, 6, '外勤工作组', '王芳', 'member');

-- 插入任务数据
INSERT INTO tasks (task_name, description, group_id, start_time, end_time, latitude, longitude, wifi, nfc) VALUES
-- 编译原理课程任务
('编译原理课程签到', '编译原理课程签到任务', 1, '2024-03-20 08:00:00', '2024-03-20 09:40:00', 30.5728, 104.0668, 1, 0),
('编译原理实验签到', '编译原理实验课程签到', 1, '2024-03-22 14:00:00', '2024-03-22 17:00:00', 30.5728, 104.0668, 1, 0),
-- 实训项目任务
('项目启动会议', '实训项目启动会议签到', 2, '2024-03-21 09:00:00', '2024-03-21 10:30:00', 30.5728, 104.0668, 1, 1),
('需求分析会议', '项目需求分析会议签到', 2, '2024-03-23 14:00:00', '2024-03-23 16:00:00', 30.5728, 104.0668, 1, 1),
-- 外勤工作任务
('客户拜访', '拜访重要客户', 3, '2024-03-22 10:00:00', '2024-03-22 12:00:00', 30.5728, 104.0668, 0, 0),
('项目现场考察', '项目现场考察签到', 3, '2024-03-24 09:00:00', '2024-03-24 17:00:00', 30.5728, 104.0668, 0, 0);

-- 插入签到记录数据
INSERT INTO sign_in_records (id, user_id, group_id, sign_in_time, sign_in_type, location, remark) VALUES
-- 编译原理课程签到记录
(1, 1, 1, '2024-03-20 08:05:00', 'normal', '教学楼A区', '正常签到'),
(2, 2, 1, '2024-03-20 08:10:00', 'normal', '教学楼A区', '正常签到'),
(3, 3, 1, '2024-03-20 08:15:00', 'late', '教学楼A区', '迟到签到'),
-- 编译原理实验签到记录
(4, 1, 1, '2024-03-22 14:05:00', 'normal', '实验楼B区', '正常签到'),
(5, 2, 1, '2024-03-22 14:10:00', 'normal', '实验楼B区', '正常签到'),
(6, 3, 1, '2024-03-22 14:20:00', 'late', '实验楼B区', '迟到签到'),
-- 项目启动会议签到记录
(7, 1, 2, '2024-03-21 08:55:00', 'normal', '会议室A', '正常签到'),
(8, 2, 2, '2024-03-21 09:00:00', 'normal', '会议室A', '正常签到'),
(9, 4, 2, '2024-03-21 09:05:00', 'normal', '会议室A', '正常签到'),
(10, 5, 2, '2024-03-21 09:10:00', 'normal', '会议室A', '正常签到'),
-- 需求分析会议签到记录
(11, 1, 2, '2024-03-23 13:55:00', 'normal', '会议室B', '正常签到'),
(12, 2, 2, '2024-03-23 14:00:00', 'normal', '会议室B', '正常签到'),
(13, 4, 2, '2024-03-23 14:05:00', 'normal', '会议室B', '正常签到'),
(14, 5, 2, '2024-03-23 14:10:00', 'normal', '会议室B', '正常签到'),
-- 客户拜访签到记录
(15, 3, 3, '2024-03-22 10:05:00', 'normal', '客户公司', '正常签到'),
(16, 4, 3, '2024-03-22 10:10:00', 'normal', '客户公司', '正常签到'),
(17, 6, 3, '2024-03-22 10:15:00', 'normal', '客户公司', '正常签到'),
-- 项目现场考察签到记录
(18, 3, 3, '2024-03-24 09:05:00', 'normal', '项目现场', '正常签到'),
(19, 4, 3, '2024-03-24 09:10:00', 'normal', '项目现场', '正常签到'),
(20, 6, 3, '2024-03-24 09:15:00', 'normal', '项目现场', '正常签到');

-- 插入任务记录数据
INSERT INTO tasks_record (task_id, user_id, username, group_id, group_name, task_name, status, signed_time, latitude, longitude, face_data, ssid, bssid, tagid, tagname) VALUES
-- 编译原理课程任务记录
(1, 1, '游辰昊', 1, '编译原理课程组', '编译原理课程签到', 1, '2024-03-20 08:05:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(1, 2, '曹航语', 1, '编译原理课程组', '编译原理课程签到', 1, '2024-03-20 08:10:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(1, 3, '刘胜杰', 1, '编译原理课程组', '编译原理课程签到', 1, '2024-03-20 08:15:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
-- 编译原理实验任务记录
(2, 1, '游辰昊', 1, '编译原理课程组', '编译原理实验签到', 1, '2024-03-22 14:05:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(2, 2, '曹航语', 1, '编译原理课程组', '编译原理实验签到', 1, '2024-03-22 14:10:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(2, 3, '刘胜杰', 1, '编译原理课程组', '编译原理实验签到', 1, '2024-03-22 14:20:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
-- 项目启动会议任务记录
(3, 1, '游辰昊', 2, '实训项目组', '项目启动会议', 1, '2024-03-21 08:55:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(3, 2, '曹航语', 2, '实训项目组', '项目启动会议', 1, '2024-03-21 09:00:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(3, 4, '张明', 2, '实训项目组', '项目启动会议', 1, '2024-03-21 09:05:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(3, 5, '李华', 2, '实训项目组', '项目启动会议', 1, '2024-03-21 09:10:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
-- 需求分析会议任务记录
(4, 1, '游辰昊', 2, '实训项目组', '需求分析会议', 1, '2024-03-23 13:55:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(4, 2, '曹航语', 2, '实训项目组', '需求分析会议', 1, '2024-03-23 14:00:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(4, 4, '张明', 2, '实训项目组', '需求分析会议', 1, '2024-03-23 14:05:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(4, 5, '李华', 2, '实训项目组', '需求分析会议', 1, '2024-03-23 14:10:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
-- 客户拜访任务记录
(5, 3, '刘胜杰', 3, '外勤工作组', '客户拜访', 1, '2024-03-22 10:05:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(5, 4, '张明', 3, '外勤工作组', '客户拜访', 1, '2024-03-22 10:10:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(5, 6, '王芳', 3, '外勤工作组', '客户拜访', 1, '2024-03-22 10:15:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
-- 项目现场考察任务记录
(6, 3, '刘胜杰', 3, '外勤工作组', '项目现场考察', 1, '2024-03-24 09:05:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(6, 4, '张明', 3, '外勤工作组', '项目现场考察', 1, '2024-03-24 09:10:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(6, 6, '王芳', 3, '外勤工作组', '项目现场考察', 1, '2024-03-24 09:15:00', NULL, NULL, NULL, NULL, NULL, NULL, NULL);

-- 插入补签申请数据
INSERT INTO check_application (group_id, task_id, task_name, user_id, username, reason, status, admin_id, admin_username) VALUES
-- 编译原理课程补签申请
(1, 1, '编译原理课程签到', 3, '刘胜杰', '因交通拥堵导致迟到，申请补签', 'approved', 1, '游辰昊'),
-- 编译原理实验补签申请
(1, 2, '编译原理实验签到', 3, '刘胜杰', '因设备故障导致迟到，申请补签', 'approved', 1, '游辰昊'),
-- 项目启动会议补签申请
(2, 3, '项目启动会议', 5, '李华', '因临时会议冲突导致迟到，申请补签', 'approved', 2, '曹航语'),
-- 需求分析会议补签申请
(2, 4, '需求分析会议', 4, '张明', '因系统故障导致无法及时签到，申请补签', 'approved', 2, '曹航语'),
-- 客户拜访补签申请
(3, 5, '客户拜访', 6, '王芳', '因客户临时改期导致迟到，申请补签', 'approved', 3, '刘胜杰'),
-- 项目现场考察补签申请
(3, 6, '项目现场考察', 4, '张明', '因天气原因导致迟到，申请补签', 'approved', 3, '刘胜杰');

-- 插入加入申请数据
INSERT INTO join_application (group_id, user_id, username, reason, status) VALUES
-- 编译原理课程组加入申请
(1, 4, '张明', '对编译原理感兴趣，希望加入学习小组', 'accepted'),
(1, 5, '李华', '希望提高编译原理知识水平', 'accepted'),
-- 实训项目组加入申请
(2, 3, '刘胜杰', '有相关项目经验，希望参与实训项目', 'accepted'),
(2, 6, '王芳', '希望积累项目经验', 'accepted'),
-- 外勤工作组加入申请
(3, 1, '游辰昊', '有外勤工作经验，希望加入外勤组', 'accepted'),
(3, 2, '曹航语', '希望积累外勤工作经验', 'accepted'); 