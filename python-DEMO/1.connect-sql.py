
import pymysql
import sys
import random, string


class databaseClass:
    conn = pymysql.connect(host = 'localhost',user = "root",passwd = "0",db = "sql_test")
    #print(conn)
    #print(type(conn))
    cur = conn.cursor()

    def createDatabase(self):
        # 打开数据库连接,不需要指定数据库,应为数据库需要创建
        # 创建数据库
        self.cur.execute('CREATE DATABASE IF NOT EXISTS pythonDb4 DEFAULT CHARSET utf8 COLLATE utf8_general_ci;')

    def closeDatabase(self):
        self.cur.close()
        self.conn.commit()  # 增删改都要要提交数据
        self.conn.close()

    def createTable(self):
        # 如果是数据库中有这个表，先drop掉，然后create表，然后再进行数据插入
        self.cur.execute('drop table if exists user')
        sql = """CREATE TABLE IF NOT EXISTS `user` (
             `id` int(11) NOT NULL AUTO_INCREMENT,
             `name` varchar(255) NOT NULL,
             `age` int(11) NOT NULL,
             PRIMARY KEY (`id`)
           ) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=0"""

        self.cur.execute(sql)

    def insertOne(self):
        #insert = self.cur.execute("insert into user values(5, 'tom', 18)")
        #print('添加语句受影响的行数', insert)
        # 另一种写法
        sql = "insert into user values(%s, %s, %s)"
        #
        try:
            # 执行sql语句
            self.cur.execute(sql, (8, '小明', 30))
            # 提交到数据库执行
            self.conn.commit()
        except:
            # Rollback in case there is any error
            self.conn.rollback()
            print("函数insertOne()出现错误:", sys.exc_info()[0])

    def insertMany(self):
        # 另一种插入数据的方式, 通过字符串传入值
        sql = "insert into user values(%s, %s, %s)"
        try:
            insert = self.cur.executemany(sql,
                                          [(4, '文明', 20), (10, '法制', 18), (6, '礼让', 19), (7, '和谐', 20), (13, '富强', 18),
                                           (9, '礼让', 19)])
            print('批量插入返回受影响的行:', insert)

        except:
            # Rollback in case there is any error
            self.conn.rollback()
            print("函数insertMany()出现错误:", sys.exc_info()[0])



    """
    查询：
    使用execute()函数得到的只是受影响的行数，并不能真正拿到查询的内容。cursor对象还提供了3种提取数据的方法：
    fetchone、fetchmany、fetchall.。每个方法都会导致游标>>>>>>>>>>>>>>移动,三个方法都必须和execute一起使用,并且在execute之前。
    cursor.fetchone():获取游标所在处的一行数据，返回元组，没有返回None
    cursor.fetchmany(size):接受size行返回结果行。如果size大于返回的结果行的数量，则会返回cursor.arraysize条数据。
    cursor. fetchall():接收全部的返回结果行。
    """


    def searchOne(self):
        self.cur.execute("select * from user;")

        while 1:
            res = self.cur.fetchone()
            if res is None:
                # 表示已经取完结果集
                break
            print(res)

    """
    注意：从execute()函数的查询结果中取数据，以元组的形式返回游标所在处的一条数据，如果游标所在处没有数据，将返回空元组，
    该数据执行一次，游标向下移动一个位置。
    fetchone()函数必须跟exceute()函数结合使用，并且在exceute()函数之后使用
    """

    def searchMany(self):
        self.cur.execute("select * from user")

        # 取3条数据
        resTup = self.cur.fetchmany(5)
        # print(type(resTup))  tuple
        for res in resTup:
            print(res)

    """
    注意:但如果游标所在处没有数据，将返回空元组,数据库中有6条数据,要查询9条也不报错,会返回6条。查询几条数据，游标将会向下移动几个位置。
    fetmany()函数必须跟exceute()函数结合使用，并且在exceute()函数之后使用
    """

    def searchAll(self):
        self.cur.execute("select * from user")
        resTup = self.cur.fetchall()
        print(type(resTup))
        print('共%d条数据' % len(resTup))

    """
    获取游标开始及以下所有的数据,并以元组的形式返回，
    游标所在处没有数据将返回空元组,执行完这个方法后,游标将移动到数据库表的最后
    """

    # 更新一条数据
    def updateOne(self):
        # 更新数据
        update = self.cur.execute("update user set age=100 where name='礼让'")
        print('修改后受影响的行数:', update)

        # 查询数据
        self.cur.execute("select * from user where name='礼让'")
        # print(self.cur.fetchall())   # ((6, '礼让', 100), (9, '礼让', 100))
        print(self.cur.fetchmany(6))  # ((6, '礼让', 100), (9, '礼让', 100))
        #  fetchmany()里面要传入数,才能返回多条.不传就返回一条
        # fetchall() 就不用传数

    # 更新多条数据
    def updateMany(self):
        # 更新前查询所有数据
        self.cur.execute("select * from user where name in ('礼让', '和谐', '富强');")
        for res in self.cur.fetchall():
            print(res)
        print('*' * 40)

        # 更新两条数据
        sql = "update user set age=%s where name=%s"
        try:
            update = self.cur.executemany(sql, [(15, '和谐'), (18, '富强')])

        except:
            # Rollback in case there is any error
            self.conn.rollback()
            print("函数updateMany()出现错误:", sys.exc_info()[0])


        # 更新2条数据后查询所有数据
        self.cur.execute("select * from user where name in ('和谐', '富强');")
        for res in self.cur.fetchall():
            print(res)

    #     删除单条数据
    def deleteOne(self):
        self.cur.execute("select * from user;")
        for res in self.cur.fetchall():
            print(res)

        # 删除某条数据
        self.cur.execute("delete from user where id='4';")
        print('*' * 40)
        self.cur.execute("select * from user;")
        for res in self.cur.fetchall():
            print(res)

    #     删除多条数据
    def deleteMany(self):
        self.cur.execute("select * from user;")
        for res in self.cur.fetchall():
            print(res)

        #         删除多条数据
        sql = "delete from user where id=%s"
        self.cur.executemany(sql, [(5), (6)])

        print('*' * 40)
        self.cur.execute("select * from user")
        for res in self.cur.fetchall():
            print(res)

    # 十五回滚
    def rollBack(self):
        # 更新前查询
        self.cur.execute("select * from user;")
        for res in self.cur.fetchall():
            print(res)

        # 更新
        sql = "update user set age=%s where name=%s"
        self.cur.executemany(sql, [('2', '和谐'), ('1', '礼让')])

        # 更新后查询下
        print('*' * 40)
        self.cur.execute("select * from user")
        for res in self.cur.fetchall():
            print(res)

        # 回滚事务
        print('*' * 40)
        print('回滚事务后的数据为：')
        self.conn.rollback()
        self.cur.execute("select * from user;")
        for res in self.cur.fetchall():
            print(res)

    #   插入100条数据---一次插入100条
    def insert100Many(self):
        words = list(string.ascii_letters)  # 生成字母
        random.shuffle(words)  # 打乱顺序

        sql = "insert into user values(%s, %s, %s)"
        for i in range(100, 200):
            self.cur.executemany(sql, [(i + 1, "".join(words[:5]), random.randint(0, 80))])

        # 插入100条数据之后查看
        self.cur.execute("select * from user")
        for res in self.cur.fetchall():
            print(res)


if __name__ == '__main__':
    # databaseClass().createDatabase()

    #databaseClass().createTable()
    databaseClass().insertOne()
    databaseClass().insertMany()
    databaseClass().searchOne()
    # databaseClass().searchMany()
    # databaseClass().searchAll()
    # databaseClass().updateOne()
    # databaseClass().updateMany()
    # databaseClass().deleteOne()
    # databaseClass().deleteMany()
    # databaseClass().rollBack()
    #databaseClass().insert100Many()


    # 关闭数据库链接
    databaseClass().closeDatabase()




    '''
    1.加上try-catch错误处理，比如插入记录时，已经存在ID为5的数据
    '''

