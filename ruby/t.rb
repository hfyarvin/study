# -*- coding: UTF-8 -*-
$LOAD_PATH << '.'
require "support"
require "tempfile"

BEGIN{
	puts "begin....=========="
}

END{
	puts "end.=========="
}
#自动换行
puts "Helllo Ruby"

puts "Helllo Ruby"

print <<EOF
 there is "here document"
 "here"
EOF

=begin
  "这是注释"	
=end

(10..15).each do |n|
	print n, ' '
end

puts ""

(1...5).each do |n|
	print n, ' '
end
puts ""

=begin
	"类"
=end

class Sample
	# 方法名总是以小写字母开头
	def hello
		puts "Hello Ruby!"
	end
end

obj = Sample.new
obj.hello

class Customer
	@@no_of_customers = 0

	def initialize(id, name, addr)
		@cust_id = id
		@cust_name = name
		@cust_addr = addr
	end

	def display_details()
		puts "Customer id #{@cust_id}"
		puts "Customer name #{@cust_name}"
		puts "Customer address #{@cust_addr}"
	end

	def total_no_of_customers()
		@@no_of_customers +=1
		puts "Total number of customers: #{@@no_of_customers}"
	end
end

# 创建对象
cust1=Customer.new("1", "John", "Wisdom Apartments, Ludhiya")
cust2=Customer.new("2", "Poul", "New Empire road, Khandala")
 
# 调用方法
cust1.display_details()
cust1.total_no_of_customers()
cust2.display_details()
cust2.total_no_of_customers()

class Example
   VAR1 = 100
   VAR2 = 200
   def show
       puts "第一个常量的值为 #{VAR1}"
       puts "第二个常量的值为 #{VAR2}"
   end
end
 
# 创建对象
object=Example.new()
object.show

=begin
	self: 当前方法的接收器对象。
	true: 代表 true 的值。
	false: 代表 false 的值。
	nil: 代表 undefined 的值。
	__FILE__: 当前源文件的名称。
	__LINE__: 当前行在源文件中的编号。
=end

puts "源文件：#{__FILE__}"

foo = false
bar = true
quu = false
 
case
when foo then puts 'foo is true'
when bar then puts 'bar is true'
when quu then puts 'quu is true'
end

=begin
	块
=end

def test
	yield 5
	puts "在test方法内"
	yield 100
end

test {|i| puts "你在块#{i}内"}

class Decade
include Week
	no_of_yrs = 10
	def no_of_months
		puts Week::FIRST_DAY #常量
		number = 10 * 12
		puts number
	end
end

d1 = Decade.new
puts Week::FIRST_DAY
Week.weeks_in_month # 方法
Week.weeks_in_year
d1.no_of_months

names = Array.new(4, "mac")
puts names.size
puts names.length
puts "#{names}"

nums = Array.new(10) { |i|  i = i * 2}
puts "#{nums}"

digits = Array(0..9)
n = digits.at(6)
puts "#{digits}"
puts "#{n}"

months = Hash.new("month")#设置默认值
puts "month0:#{months[0]}"
puts "month72:#{months[72]}"

H = Hash["a" => 100, "b" => 120]
puts "#{H["a"]}"
puts "#{H["b"]}"
keys = H.keys
puts "keys: #{keys}"

=begin 
	日期
=end

time = Time.new
puts "当前时间：" + time.inspect
time2 = Time.now
puts "当前时间：" + time2.inspect
puts "当前时间 : " + time.inspect
puts time.year    # => 日期的年份
puts time.month   # => 日期的月份（1 到 12）
puts time.day     # => 一个月中的第几天（1 到 31）
puts time.wday    # => 一周中的星期几（0 是星期日）
puts time.yday    # => 365：一年中的第几天
puts time.hour    # => 23：24 小时制
puts time.min     # => 59
puts time.sec     # => 59
puts time.usec    # => 999999：微秒
puts time.zone    # => "UTC"：时区名称


# 返回从纪元以来的秒数
timeunixs = Time.now.to_i  
 
# 把秒数转换为 Time 对象
Time.at(timeunixs)
 
# 返回从纪元以来的秒数，包含微妙
timeformnow = Time.now.to_f

 
a = [1,2,3,4,5]
b = a.collect{|x| 10 * x}
puts b

sum = 0 
cutcome = {"block1" => 1000, "book2" => 1000, "book3" => 4000}
cutcome.each{|item, price| sum += price}
print "sum = " + sum.to_s

puts ""

sum = 0
cutcome = {"block1" => 1000, "book2" => 1000, "book3" => 4000}
cutcome.each{|pair| sum += pair[1]}
print "sum = " + sum.to_s

puts ""

puts "Enter a value :"
# enter_val = gets
# puts enter_val

str = "Hell Arvin"
putc str
puts ""


=begin
	--RANGE--
	('a'..'d')    #==> 'a', 'b', 'c', 'd'
=end

=begin
	FILE
=end
f = File.new("support.rb", "r")

if f
	content = f.sysread(20)
	puts content
elsif
	puts "Unable to open file!" 
end

p "------------------"
IO.foreach("support.rb"){|block| puts block}
p "------------------"

p "------------------"
f2 = Tempfile.new('tingtong')
f2.puts "Hello"
puts f2.path
f2.close
p "------------------"


=begin
	--RANGE--
	('a'..'d')    #==> 'a', 'b', 'c', 'd'
=end

range2 = ('baa'..'bac').to_a
puts "#{range2}"

# each
ary = [1 ,2 ,3 ,4 ,5]
ary.each do |i|
	puts i
end
puts "----------"
# collect
ary2 = ary.collect{|i| i*3}
puts ary2

sum1 = 0
sum2 = 0
cutcome = {"block1" => 1000, "book2" => 1000, "book3" => 4000}
cutcome.each{|item, price| sum1 += price}
puts "sum1 = " + sum1.to_s

cutcome.each{|pair| sum2 += pair[1]}
puts "sum2 = " + sum2.to_s

=begin
	CLASS
	attr_accessor
	attr_reader
	attr_writer
=end

class Box
	def initialize(width,height)
		@width, @height = width,height
	end

	def printWidth
		@width
	end

	def printHeight
		@height
	end
	
	# setter
	def setWidth(v)
		@width = v
	end

	def setHeigth(v)
		@height = v
	end
end

b = Box.new(10, 20)
b.setWidth(5)
b.setHeigth(6)
x = b.printWidth()
y = b.printHeight()


puts "H: #{y}"
}
puts "W: #{x}"