### Android自动化测试工具 UiAutomator与Espresso基本使用

- 写UI自动化测试用例，归结起来就是3步：

定位View控件

操作View控件

校验View控件的状态

####UiAutomator

- UiDevice：

设备对象，通过UiDevice的getInstance(instrumentation)方法获取，可以通过UiDevice实例来检测设备的各种属性，比如获取屏幕的方向、尺寸等，还可以通过UiDevice实例来执行设备级别的操作，比如点击Home键、返回键等

- UiSelector

用于获取某些符合条件的UI控件对象，可以通过资源id、描述等熟悉获取

- UiObject

代表一个UI控件，通过uiDevice的findObject(UiSelector)方法获取，获取到UiObject实例后，就可以对UI控件进行相关的操作，比如点击、长按等

- UiCollection

代表UI控件集合，相当于ViewGroup，比如界面中有多个CheckBox时，可以通过类名获取到当前界面下的所有CheckBox，然后通过控件id获取指定的CheckBox对象

- UiScrollable

代表可滚动的控件

- 测试用例

```java
public class UiTest extends TestCase {

    public void testA() throws UiObjectNotFoundException {
        // 获取设备对象
        Instrumentation instrumentation = InstrumentationRegistry.getInstrumentation();
        UiDevice uiDevice = UiDevice.getInstance(instrumentation);
        // 获取上下文
        Context context = instrumentation.getContext();

        // 启动测试App
        Intent intent = context.getPackageManager().getLaunchIntentForPackage("com.yang.designsupportdemo");
        intent.addFlags(Intent.FLAG_ACTIVITY_CLEAR_TASK);
        context.startActivity(intent);

        // 打开CollapsingToolbarLayout
        String resourceId = "com.yang.designsupportdemo:id/CollapsingToolbarLayout";
        UiObject collapsingToolbarLayout = uiDevice.findObject(new UiSelector().resourceId(resourceId));
        collapsingToolbarLayout.click();

        for (int i = 0; i < 5; i++) {
            // 向上移动
            uiDevice.swipe(uiDevice.getDisplayHeight() / 2, uiDevice.getDisplayHeight(),
                    uiDevice.getDisplayHeight() / 2, uiDevice.getDisplayHeight() / 2, 10);

            // 向下移动
            uiDevice.swipe(uiDevice.getDisplayHeight() / 2, uiDevice.getDisplayHeight() / 2,
                    uiDevice.getDisplayHeight() / 2, uiDevice.getDisplayHeight(), 10);
        }

        // 点击应用返回按钮
        UiObject back = uiDevice.findObject(new UiSelector().description("Navigate up"));
        back.click();

        // 点击设备返回按钮
        uiDevice.pressBack();
    }
}
```

####Espresso

- Espresso测试框架是在Android测试支持库中的，提供了操作一个app的api来模拟用户的操作。Espresso测试可以运行在Android2.3.3及其更高版本的机器上。使用Espresso的一个优点是对测试操作提供了自动的同步操作。Espresso会在主线程空闲时检测到，会在恰当的时候执行测试命令，这提高了测试的可靠性。

- 定位元素onView

　　onView使用的是一个hamcrest匹配器，该匹配器只匹配当前视图层次结构中的一个（且只有一个）视图。如果你不熟悉hamcrest匹配器，建议先看看这个。通常情况下一个控件的id是唯一的，但是有些特定的视图是无法通过R.id拿到，因此就需要访问Activity或者Fragment的私有成员找到拥有R.id的容器。有的时候也需要使ViewMatchers来缩小定位的范围。最简单的onView就是这样的形式：onView(withId(R.id.my_view))
　　　　
- 定位元素onData

假设一个Spinner的控件，要点击“Americano”，使用默认的Adaptor，它的字段默认是String的，因此当要进行点击的时候，就可以使用如下方法： onData(allOf(is(instanceOf(String.class)),is("Americano"))).perform(click());


- perform()参数中常用的方法:

在上面的一段官网代码中,我们用到了perform(click()),那么除了click()方法还有其他功能强大的方法可以供我们使用,下面列举一些常用的方法:

click(): 
返回一个点击action,Espresso利用这个方法执行一次点击操作,就和我们自己手动点击按钮一样,只不过Espresso把点击这个操作自动化了,下面的方法都是一样的道理,就不再赘述了.

clearText(): 
返回一个清除指定view中的文本action,在测试EditText时用的比较多

swipeLeft(): 
返回一个从右往左滑动的action,这个在测试ViewPager时特别有用

swipeRight(): 
返回一个从左往右滑动的action,这个在测试ViewPager时特别有用

swipeDown(): 
返回一个从上往下滑动的action

swipeUp(): 
返回一个从下往上滑动的action

closeSoftKeyboard(): 
返回一个关闭输入键盘的action

pressBack(): 
返回一个点击手机上返回键的action

doubleClick(): 
返回一个双击action

longClick(): 
返回一个长按action

scrollTo(): 
返回一个移动action

replaceText(): 
返回一个替换文本action

openLinkWithText(): 
返回一个打开指定链接action