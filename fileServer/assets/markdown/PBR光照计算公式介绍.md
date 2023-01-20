# PBR光照计算公式介绍

参考视频：  
[图形 2.4 传统经验光照模型详解](https://www.bilibili.com/video/BV1B54y1j7zE)  
[GAMES101-现代计算机图形学入门-闫令琪 P15](https://www.bilibili.com/video/BV1X7411F744?p=15)    
参考资料：  
[PBR-learnopengl](https://learnopengl.com/PBR/Theory)  
[彻底看懂 PBR/BRDF 方程-知乎](https://zhuanlan.zhihu.com/p/158025828)  
[辐射强度、辐亮度、辐照度——一文搞定](https://blog.csdn.net/a6333230/article/details/82968484)  
[辐射照度、辐射强度、光照度、发光强度（差异以及如何相互转换）（易懂讲解）](https://blog.csdn.net/a6333230/article/details/90036993)  
[【基于物理的渲染（PBR）白皮书】（一） 开篇：PBR 核心知识体系总结与概览](https://zhuanlan.zhihu.com/p/53086060)  
[【实时渲染】菲涅尔反射率](https://blog.csdn.net/Terie/article/details/109460237)  
[PBR 以及在 Disney 和 UE 渲染模型中的使用](https://blog.csdn.net/leonwei/article/details/104044122/)  
[Adopting a physically based shading model 这篇文章提供了大量的论文参考](https://seblagarde.wordpress.com/2011/08/17/hello-world/)  
[Physically Based Shading at Disney](https://blog.selfshadow.com/publications/s2012-shading-course/burley/s2012_pbs_disney_brdf_notes_v3.pdf)  
[Real Shading in Unreal Engine 4](https://de45xmedrsdbp.cloudfront.net/Resources/files/2013SiggraphPresentationsNotes-26915738.pdf)

## 能量守恒在光照模型中的作用

PBR（基于物理的渲染）是对现实近似模拟的渲染技术，旨在物理上合理地模拟光照效果，因此效果往往比非 pbr 渲染要好
PBR 模型前置理论基础：微表面理论，能量守恒，菲涅尔反射

## 微表面理论

微表面理论认为任何平面都是由极小的“微平面组成”
<img src="https://c2.im5i.com/2023/01/10/YP0o5.md.png" width="80%">
在微观层面上没有任何表面是完全光滑的，但鉴于这些微表面足够小，我们无法在每个像素的基础上区分它们，我们用一个“粗糙度”参数来估计微表面的粗糙度。再根据微表面的粗糙度，我们可以计算出微平面法线 m 与向量 h 一样的微平面的比例（这个比例具体如何计算后文会解释），这个 h（halfway vector）也是 blinn-phong 模型使用到的半角向量，h 的计算方式如下图（l:表面指向光源方向 v:表面指向视线方向）
<img src="https://c2.im5i.com/2023/01/10/YPrzZ.png">
从下图中可以看到只有法线等于 h 的微平面才能使光线反射到视线上。因此只要知道出与向量 h 一样的微平面的比例，就可以知道有多少微平面对观察者观察到的颜色有贡献

<img src="https://c2.im5i.com/2023/01/10/YPYJs.md.png">

粗糙度越大镜面反射范围越大，更暗；越小反射范围越小，越亮
<img src="https://c2.im5i.com/2023/01/10/YP7Dw.png" width="80%">
微平面理论近似的遵循能量守恒定律：出射光的能量不能大于入射光的能量。（从上图也能看出，高光范围小更亮，范围大更暗）所以遵循能量守恒定律是为了让场景看起来更符合物理，更真实

微平面理论说明光线在交界点处会在多个方向上产生反射和折射光。
<img src=" https://c2.im5i.com/2023/01/10/YP5VD.md.png" width="60%">

## 菲涅尔反射

想象观察一个湖面，远看（入射角大）树的倒影（反射）很清晰，近看（入射角小）湖面下的鱼（折射）清晰。可见<font color=#ff0000>随着入射角增大（从近到远看），光的反射率单调上升，折射率单调下降</font>。更容易看到物体表面上其他物体的倒影而很难看到物体内部的情况

<font color=#ff0000>菲涅尔公式则描述了一束光经过两个介质交界面分裂成反射光和折射光时，反射光与折射光占原光线的比例</font>。（图中可正可负的意思是正负号与如何定义光线的正方向有关）
值得注意的是，折射光进入介质内部后可能会再次发生反射和折射，这部分被称为“次表面散射”，这些光部分会被吸收，部分会再次离开物体表面散射出去，形成**漫反射**。着色器用次表面散射可以以牺牲性能为代价显著改善皮肤，蜡或大理石等材料的视觉效果
<img src="https://c2.im5i.com/2023/01/10/YVQwn.png" width="40%">
<img src="https://c2.im5i.com/2023/01/10/YPtMP.md.png" width="60%">
rs 是反射光的垂直分量，rp 是反射光的水平分量，ts 是折射光的垂直分量，tp 是折射光的水平分量

对于自然光，s 波和 p 波的能量相等，因此自然光的反射率<img src="https://c2.im5i.com/2023/01/20/Y5tqz.png" alt="Y5tqz" width="598" height="90" data-is360="0" data-load="full" class="" style="width: 598px; height: 90px;">
当光线入射角趋近于 0°（垂直于介质表面）时，有（n 为折射率）
<img src="https://c2.im5i.com/2023/01/20/Y5uE5.png" alt="Y5uE5" width="301" height="86" data-is360="0" class="" style="width: 301px; height: 86px; display: block;">
此时<img src="https://c2.im5i.com/2023/01/20/Y5ULG.png" alt="Y5ULG" width="343" height="87" data-is360="0" class="" style="width: 343px; height: 87px; display: block;">
<font color=#ff0000>接近 0° 入射角时的菲涅尔反射率 Rn（表示反射光占原光线的比例）被记作 F0</font>我们可以用上式，根据折射率（也称 IOR）计算 F0
不同材质其 F0 不同（线性和 sRGB 指色彩空间）,如下图所示

<img src="https://img-blog.csdnimg.cn/21ad478720114d52852c339ea4897f59.png" width="80%">

下图展示了不同波长和不同入射角下玻璃，铜和铝的菲涅尔反射率



<img src="https://c2.im5i.com/2023/01/10/YVU84.md.png" width="80%" >

结合以上两图我们可以发现金属和非金属的菲涅尔反射有很大差异：

### 金属

<img src="https://img-blog.csdnimg.cn/c27e07451cce479eaf278940a8d50f0c.png" width="60%">
·菲涅尔反射率会受到波长影响，需要引入复数来表示反射率，因此用Fresnel-Schlick近似(下文会介绍)时需要用rgb三个值（且这三个值不同，体现对不同光有不同的吸收率）表示
·反射率大部分在0.5和1之间，说明金属吸收了大部分折射光，其颜色主要是由镜面反射决定的

### 非金属（电介质）

<img src="https://img-blog.csdnimg.cn/2ee40cbb2b2e42f4b69314fbb41b9bfc.png" width="60%">

·菲涅尔反射率与波长无关，因此 F0 只用一个值（一般取 0.04）就能表示
·反射率几乎接近 0，说明非金属的颜色主要是由漫反射（也就是折射光引起的次表面散射）决定的
（注意这里的漫反射和镜面反射是从微观角度上考虑的，看光线是反射光还是次表面散射光来区分，而平时说的漫反射和镜面反射从宏观上观察物体是否光滑来区分）

<img src="https://img-blog.csdnimg.cn/b672ebe899cf40608024f759b5de3986.png" width="60%">

也许有人就会产生疑惑：根据微平面理论，直接反射光方向是多种多样的，那为什么这部分光不算漫反射而算镜面反射呢？ 1.漫反射和镜面反射本质上都是光线，我们平时说的漫反射和镜面反射是从宏观角度观察物体，看物体表面是否光滑（光线是否集中），也就是根据结果来决定光线是镜面反射还是漫反射；而在 PRB 模型中，从微观角度观察光线，看光线是直接反射，还是次表面散射，也就是根据光线的路径决定光线是镜面反射还是漫反射 2.结合实际代码来理解，我们在像素着色器中计算 BRDF，这里一个像素可以看作一个微元（一个微元有许多微平面），且拥有唯一一个法线，这个法线就代表这所有微平面法线的平均值，也就是说大体上微平面的法线都是与微元的法线一致，因此直接反射光主要集中在对称方向，且集中程度与粗糙度有关（可以看到当粗糙度变大，镜面反射的结果实际上就是我们从宏观角度所说的“漫反射”），所以用镜面反射描述直接反射光是比较精确的

## 反射比方程/渲染方程(reflectance equation)

该方程是目前模拟光的视觉效果的最佳模型，具体如下
<img src="https://img-blog.csdnimg.cn/1e484b34b9984c308ff78eda93c3c522.png" width="80%">
该方程描述了一个极小的微元表面接受光照后发出的光的功率。让我们先来看看这些字母分别代表什么

p:辐射功率。<font color=#ff0000>指单位时间内辐射源所发射的总辐射能</font>
单位 瓦

ω:<font color=#ff0000>ωi 表示入射光线（l） ，ω0 表示出射光线（v）</font>，注意等号右边要积分的微元指的是立体角

dω:立体角，是站在某一点的观察者测量到的物体大小的尺度
锥体的立体角大小定义为：以锥体的顶点为球心作球面，该锥体在球表面截取的面积与球半径平方之比，
单位 球面度（一整个球的球面度为 4π）
计算方式见下图
<img src="https://img-blog.csdnimg.cn/e4e8d5654fb5437ba8999e05f0fce479.png" width="80%">
&nbsp;

L:辐射亮度(radiance)，简称辐亮度。<font color=#ff0000>指面辐射源在单位时间内通过垂直于给定方向的平面上单位面积、单位立体角上辐射出的能量</font>，即辐射源在单位投影面积上、单位立体角内的辐射通量
计算方式：L=ddΦ/(dAdΩ·cosθ)，θ 为立体角(solid angle)与法线之间的夹角
单位 瓦/(球面度·米^2)
辐射亮度表示面辐射源上某点在一定方向上的辐射强弱的物理量
<img src="https://img-blog.csdnimg.cn/441cf184cf99497a88cfb8c46bfb62d8.png" width="40%">
&nbsp;

Φ:辐射通量。<font color=#ff0000>指单位时间内通过某一截面的辐射能</font>
单位 瓦(焦耳/秒)

E:辐射照度(Irradiance)，简称辐照度。<font color=#ff0000>指接收物体的单位表面积上接收到的辐射功率</font>
单位 瓦/米^2
计算方式见下图
<img src="https://img-blog.csdnimg.cn/c62ad93c32724ecfa2d2f23196120204.png" width="70%">
&nbsp;

下图直观地展示了这些物理量的区别和联系
<img src="https://img-blog.csdnimg.cn/38e2e5b4c1b74bf4a25578a410b12dd6.png" width="80%">
<img src="https://img-blog.csdnimg.cn/721bffbbf79b459d8fe94d86d15916df.png" width="80%">
（注意这里的辐亮度计算中分母的 dA 已经是投影面积，与上面的微元计算公式 L=ddΦ/(dAdΩ·cosθ)不同，后者的 da 是光源上的微元面积因此要乘 cosθ，以将其转换到垂直于光线的平面，即与上图介绍辐射亮度的图中黄色平面平行)

了解了这些字母的含义之后我们再回看渲染方程
<img src="https://img-blog.csdnimg.cn/1e484b34b9984c308ff78eda93c3c522.png" width="80%">
该方程本质上告诉我们，给定在物体（物体接受光照后发出光也能看作辐射源）上一点 A 和入射光的信息，就能计算出射光的辐亮度(特指被眼睛观察到的出射光的辐亮度)
·L0 表示 A 点出射光的辐亮度
·fr 表示出射光辐亮度和入射光辐照度的比例函数
fr 有多种形式的函数可以模拟，大体分为如下几种
1.BRDF（双向反射分布函数）：仅处理受光面，且不考虑次表面散射（因为漫反射属于次表面散射，所以严格地讲只是简单地考虑次表面散射）。适合不透明材质。本文只介绍 BRDF
2.BTDF（双向透射分布函数）：仅处理背光面，且不考虑次表面散射。
3.BSDF（双向散射【反射+透射】分布函数）：处理受光面和背光面，且不考虑次表面散射。适合透明度比较高的材质。
4.BSSRDF（双向散射表面反射率分布函数）：处理受光面和背光面，考虑次表面散射，适合半透明材质，云，玉石，牛奶等。

·ωi 表示入射光线光线方向 l），w0 表示出射光线（特指视线方向 v）
·Li 表示 A 点入射光的辐亮度，与 fr 函数相乘，再乘上 n·ωi 即可得到出射光的辐亮度 L0
·n·ωi 指入射光线和法线之间的夹角的余弦值，用于将入射光的辐亮度 Li 转换为辐照度 E
·∫Ω dωi 则是入射光半球的积分（因为光不可能同时照亮微元的两个面，所以是半球而不是整个球，这里的光并没有指定光源的种类，只要是从微元的正半球照射进来的光线都要考虑），可以理解为对入射半球上所有的 ωi 进行累加
fr 的推导见[彻底看懂 PBR/BRDF 方程 2. BRDF 到底是啥?](https://zhuanlan.zhihu.com/p/158025828)
接下来讨论怎么具体表示 BRDF（fr 函数）
BRDF 分为 diffuse（漫反射）和 specular（镜面反射）两项

### diffuse BRDF（漫反射）

<img src="https://img-blog.csdnimg.cn/4ceb6fec84904afa83575a9073b0418c.png">
漫反射的计算方法大体分为两种，基于经验（如Lambert（朗伯/兰伯特））和基于物理。

1.Lambert diffuse
Lambert diffuse 适用于朗伯辐射体：辐射源各方向上的辐射亮度不变，辐射强度随观察方向与面源法线之间的夹角 θ 的变化遵守余弦规律。
表达式如下:

<img src="https://c2.im5i.com/2023/01/20/Y5EHW.png">
（Cdiffuse 是材质本身的颜色，也被称作 albedo（反射率）/basecolor（基础色）/surfaceColor（表面颜色））
正向推导：
<img src="https://img-blog.csdnimg.cn/e708030a4e334612a24a82d8f1a639b3.png" width="80%">
可能会有人感到疑惑：这个 Lambertian 好像和我们说的兰伯特漫反射（n·l）不一样，但是我们令 fr=c/π 代回到反射比方程中，可以发现积分后 π 就被抵消掉了，而且乘上后面的 n·ωi 正好就是传统的 lambert 漫反射光照模型。因此我们可以把分母看作一个为了满足能量守恒定律而引入的归一化因子

2.Disney diffuse
<img src="https://img-blog.csdnimg.cn/53e1097f074f45c29cea9086ff261d70.png" width="80%">
FD90 指的是法线和视线呈 90° 时的菲涅尔反射率

### specular BRDF（镜面反射）

<img src="https://img-blog.csdnimg.cn/9b6e75e00a1241dd95b098c2e11ad447.png">

目前业界广泛采用的基于微表面理论的[Microfacet Cook-Torrance BRDF 模型](https://www.cs.cornell.edu/~srm/publications/EGSR07-btdf.pdf)，公式如下（l:表面指向光源方向，v:表面指向视线方向，h:半角向量（微平面法线），n:微元法线）：
<img src="https://img-blog.csdnimg.cn/847af04c43d548558042e823267f6459.png">

（推导见[彻底看懂 PBR/BRDF 方程-知乎 7.镜面反射的 BRDF 如何推导？]）(https://zhuanlan.zhihu.com/p/158025828)

可以看到该公式由分子的三个符号和分母的一个归一化因子组成，其中分子的 DFG 各代表微平面表面特性的一个近似描述的函数

#### D：Normal distribution function（NDF）法线分布函数

<font color=#ff0000>法线分布函数表示法线与半角向量方向相同的点微平面占微元的比例</font>
目前常用的是 Disney 的 Trowbridge-Reitz（各向同性的 GGX）模型，其中 α 是控制参数，这里 α 通常表示 roughness^2（由 roughness 粗糙度对 α 进行映射，当然令 α=roughness 也没问题，毕竟 roughness 是手动控制的参数）（粗糙度在 0-1 之间，越大越粗糙）
<img src="https://img-blog.csdnimg.cn/f70dc27965824300988f3ff146483c47.png" alt="在这里插入图片描述">
GGX 优点：1.成本低廉 2.更长的拖尾显得更自然
<img src="https://img-blog.csdnimg.cn/e382d73be36244b086547563b2d1a126.png">

其他 NDF 模型（m 是微平面法线，这里指半角向量 h）
<img src="https://img-blog.csdnimg.cn/119f9e57ec744921bce9853093f85c5c.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAd3NXaW5k,size_20,color_FFFFFF,t_70,g_se,x_16" alt="在这里插入图片描述">

#### F：Fresnel equation 菲涅尔方程

<font color=#ff0000>菲涅尔方程表示反射到视线方向上的光占入射光的比例</font>，菲涅尔方程是一个相当复杂的方程，一般我们用 Fresnel-Schlick 近似表达菲涅尔反射率 F
<img src="https://img-blog.csdnimg.cn/f3722f9bf4ff405e8b99e69fb2af53cf.png" alt="在这里插入图片描述">
其中 h 是半角向量（也就是微平面的法线），v 是视线方向矢量,h·v 也说明视线离半角向量越近，菲涅尔反射越强,F0 是接近 0° 入射角时的菲涅尔反射率（上文介绍过如何求 F0）
ue4 在 2013siggraph 上给出了这个形式的略微优化版本
<img src="https://img-blog.csdnimg.cn/02dd3f4f86ad416ab1a0318657caba5e.png" alt="lut">

#### G：Geometry function 几何函数（阴影遮罩函数）

<img src="https://img-blog.csdnimg.cn/f95b57890a0d493891cf4271f0d90e22.png">
之前讲微平面理论的时候提到过只有法线等于h的微平面才对brdf做出贡献,但是法线等于h的微平面中可能会有一部分发出的反射光被其他微平面给阻挡，
而<font color=#ff0000>几何函数就是描述法线等于h且未被遮蔽的微平面占微元的百分比。</font>

1.Schlick-GGX 模型
Schlick-GGX 模型是 GGX 和 Schlick-Beckmann 的的结合：
<img src="https://img-blog.csdnimg.cn/2ac6047dc8fb487eaa12dbbd3811d89c.png" alt="在这里插入图片描述">
其中 K 是对 α（注意 α 与 roughness 的转换可能因引擎而异，下图的 α 实际上就是指 roughness）的重新映射，与场景使用直接照明还是 IBL（基于图像的照明，后文介绍环境光时会介绍）有关，这里反射率方程用于描述直接光，取(α+1)^2/8
<img src="https://img-blog.csdnimg.cn/a28e20ac93e148de8e0227c24c322ad6.png" alt="在这里插入图片描述">
考虑到视线方向的几何遮挡和光方向矢量的几何阴影，我们根据 Smith 方法来包含两者得到最终的 G 值（Gsub 指的就是 Gschlick-GGX）：
<img src="https://img-blog.csdnimg.cn/fc375c8e485b41e98c01ed77a0895614.png" alt="在这里插入图片描述"><img src="https://img-blog.csdnimg.cn/a8fe12e19e144ce78d7d71f29258a78b.png">
这个模型也是 ue 在[siggraph2013 ue 第 29 幅图](https://de45xmedrsdbp.cloudfront.net/Resources/files/2013SiggraphPresentationsNotes-26915738.pdf)也提到过
<img src="https://img-blog.csdnimg.cn/31b26522b5b94929aeabca47b60f8d7e.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAd3NXaW5k,size_20,color_FFFFFF,t_70,g_se,x_16" alt="在这里插入图片描述">

2.Cook-Torrance 模型
<img src="https://img-blog.csdnimg.cn/80ba5ff3106141ba8859ec4a8a3d3b5f.png">
<img src="https://img-blog.csdnimg.cn/cb87ba4f13294879b071cdb6391bb6c4.png">
想了解更多的 diffuse 和 specular 模型计算方式可以参考这两篇
[UE4 中的基于物理的着色（一）](https://zhuanlan.zhihu.com/p/34473064)
[镜面反射 BRDF 模型(Specular BRDF)及实现效果](https://blog.csdn.net/qq_35312463/article/details/108123270)

### 将漫反射与镜面反射合并

Cook-torrance 模型提供了一种方案（据 learnopengl 所述）
<img src="https://img-blog.csdnimg.cn/6e3e4a649dac4b4e8b6295698e0ea534.png" alt="在这里插入图片描述">
<img src="https://img-blog.csdnimg.cn/d574b3a678f24b688c471d8e1c14ca0a.png" alt="在这里插入图片描述">

其中 ωi 表示入射光向量(l)，ω0 表示出射光向量(v)，kd 是漫反射系数，ks 是镜面反射系数（菲涅尔反射率），kd=1-ks（考虑到能量守恒。kd 和 ks 表示入射光照射到物体表面微元后分配给漫反射和镜面反射的比例）
考虑到 F 实际隐含了 Ks（F 表示镜面反射光占入射光的比例，所以 F=Ks），因此代入的过程中 Ks 应该删去，实际公式如下：
<img src="https://img-blog.csdnimg.cn/bad4d7d18adf4c5195b81dd982f9c627.png" alt="在这里插入图片描述">
·对于电介质，kd=1（理论上应该是取 1-F），F0 一般取（0.04，0.04，0.04）就能代表大部分电介质的特性。
·对于导体/金属，我们引入了一个 metalness（金属度）可控参数来对 F0 进行插值：  
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; **F0 = mix((0.04,0.04,0.04), albedo, metalness)**  
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; **kd=(1-ks)\*(1-metalness)**  
（可以看到 metalness=0 时 F0=0.04（等价于电介质），metalness=1 的时候则采用金属自身的反射率（F0）且没有漫反射）
(考虑到金属自由电子对光的吸收，因此 ks+kd 不一定等于 1)
<img src="https://img-blog.csdnimg.cn/a8e8a0741fd847cb87f21f00131acfe8.png" alt="在这里插入图片描述">

总结：直接光照就是上面的反射率公式，对于物体表面上的一个微元。给定微元法线正向的半球上所有的入射光的功率，我们可以得到物体在接收入射光后最终在视线方向上发出的光的功率。
在实际代码编写中，如果用反射率方程计算直接光照，环境光照另外单独计算，我们可以使用如下等价公式来简化只有直接光照的反射率方程：（中间的运算符表示对两边的矩阵或向量中的每个一 一对应的元素直接相乘，Clight 指光源颜色）
<img src="https://img-blog.csdnimg.cn/47c63aa6653446d090c8dbff2b85ea18.png" alt="在这里插入图片描述">
解析：Clight 光源颜色与波长相关，波长与能量成反比，在相同时间内，能量与功率成正比，所以这里的 L0（V）可以看作最终物体表面微元的颜色
翻译如下:
<img src="https://img-blog.csdnimg.cn/5bc91f3aa0ce42b587fdff2a5130e3f6.png">
图中（镜面反射+漫反射）还要再乘上 π（半球积分后得到），考虑到镜面反射不像漫反射那样考虑来自四面八方的光，只考虑光源发出的指向微元的光，所以不用乘上 π，漫反射乘上 π 后分母约掉
<img src="https://img-blog.csdnimg.cn/a3282d58a36141bf8a5580fa84aa9dfb.png#pic_center" alt="在这里插入图片描述">)

## 环境光

除了直接光以外我们还需要环境光（来自环境其他物体的光，默认等同于间接光），在传统经验光照模型中我们用一个常数来代替，在 PBR 中我们用反射率方程更精确地描述环境光对物体的影响（环境光与直接光很大的区别是，直接光方向是确定的，而环境光方向是来自四面八方其他物体的反射光的）
<img src="https://img-blog.csdnimg.cn/bbf044162699415abf83c590ecb74ec6.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAd3NXaW5k,size_20,color_FFFFFF,t_70,g_se,x_16" alt="在这里插入图片描述">
先介绍环境光照的表示形式
我们需要一种方式来表示来自各个方向的环境光的辐亮度（假设各个方向的环境光照都是来自无穷远处的） 1.球谐函数：[详细介绍](https://zhuanlan.zhihu.com/p/153352797) 2.球形高斯：[详细介绍](https://therealmjp.github.io/posts/sg-series-part-2-spherical-gaussians-101/)（中文翻译版本：[SG Series Part 2: 球形高斯基础](https://zhuanlan.zhihu.com/p/343751063)）
我们可以用一张图片来记录这些球形函数，这种图片叫做环境贴图 1.经纬度贴图 2.球形贴图 3.立方体贴图
（[更具体的介绍请参考这篇](https://zhuanlan.zhihu.com/p/144910975)）

了解环境光照的表示形式后我们再来看环境光影响物体的表达式

环境光的反射率方程与直接光照的一样
<img src="https://img-blog.csdnimg.cn/fa826164d576492cb703013e5d329ebb.png">
可以看到漫反射和镜面反射可以拆分，于是变成如下形式
<img src="https://img-blog.csdnimg.cn/a354fb0b25a7444b8f005833e0fe8fd1.png">
环境光的漫反射与镜面反射需要用到预计算的贴图来存储信息，因此也被称为 IBL（基于图像的光照)，接下来我们分别详细讨论。

### 环境光-漫反射

对于漫反射部分，将常数移到外面，得到如下形式：
<img src="https://img-blog.csdnimg.cn/4a209b53e9da4aeaac762ca9ff39ddeb.png">
在直接照明-漫反射中，因为入射光线是确切知道方向的，所以很容易求出射光方向的辐亮度 L0。但是环境光-漫反射中，对于每一个微元而言，来自周围环境的入射光不是唯一的。每个入射光方向 wi 都可能会有辐射，所以求解积分很复杂，这也给我们提出了两个要求： 1.需要一个方法来检索任何 wi 方向的环境光辐射度 2.求解积分必须实时快速
下面重点介绍一种解决方案--立方体贴图
**·1.立方体贴图（cube mapping）**
立方体贴图是事先预计算来自各个方向入射光的辐亮度 Li，并存储到一张立方体贴图中（一个纹素对应一个出射光方向 ω0）。
预计算方式：
对于每一个出射光 ω0，我们构建一个正向的半球，在其中离散的取大量不同方向的入射光 ωi 并采样，然后求平均，得到的结果作为出射光 ω0 的辐照度，将出射光方向作为索引，结果以颜色值的形式存储在立方体贴图中
<img src="https://img-blog.csdnimg.cn/a272d31dfc824c35abb3adb68fe08a28.png" width="70%">
因此这张立方体贴图也被称为辐照度环境贴图（irradiance map，也有把 Irradiance Environment Mapping 翻译成辉度环境映射的实在是有点难以理解）
（辐射方程还依赖于辐照度环境贴图的中心位置 p，在场景中不同的位置得到的辐照度环境贴图不一样。渲染引擎通过在整个场景中放置反射探针（reflection probes）来解决此问题，每个反射探针计算其自身周围环境的辐照度贴图。这样，对于任意位置处的辐照度就可以通过离其最近的几个反射探头的辐照度的插值来得到。现在我们假设总是从辐照度环境贴图的中心位处采样，不考虑插值的事情）
下图是立方体环境贴图和其生成的立方体辐照度贴图的实例（尽管特别像高斯模糊但并不是直接模糊处理）
<img src="https://img-blog.csdnimg.cn/8d2036f81e464dd6bb5fd3500889ae8d.png" width="80%">
具体如何生成？
参考：
我们把方程转换成如下形式：
<img src="https://img-blog.csdnimg.cn/af3670a396e84a968fd81969644803a6.jpg?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAd3NXaW5k,size_20,color_FFFFFF,t_70,g_se,x_16#pic_center" alt="在这里插入图片描述">
<img src="https://img-blog.csdnimg.cn/f7e9d7ac82e84424b300870f15ea86cf.png" alt="在这里插入图片描述">
将积分转换为离散的累加和形式，n2 对应不同的 θ，n1 对应不同的 φ
<img src="https://img-blog.csdnimg.cn/e878c52392d64f37a9007c1881ffab32.png" width="80%">
所以我们规定采样步长 dθ 和 dφ 就能对环境贴图进行采样，示例代码如下（来自[LearnOpenGL 学习笔记—PBR：IBL](https://blog.csdn.net/weixin_43803133/article/details/110385305))：

```c
============================Vertex   Shader=======================
#version 330 core
layout (location = 0) in vec3 position;
layout (location = 1) in vec3 normal;
layout (location = 2) in vec2 texcoord;
layout (location = 3) in vec3 color;

out vec3 worldPos;

uniform mat4 viewMatrix;
uniform mat4 projectMatrix;

void main(){
	worldPos = position;
	gl_Position = projectMatrix * viewMatrix * vec4(position,1.0f);
}
============================Fragment Shader=======================
#version 330 core
out vec4 FragColor;
in vec3 worldPos;

uniform samplerCube environmentMap;

const float PI = 3.14159265359;

void main()
{
	// 世界向量充当原点的切线曲面的法线，与WorldPos对齐。
	// 给定此法线，计算环境的所有传入辐射。
    vec3 N = normalize(worldPos);

    vec3 irradiance = vec3(0.0);

    // 计算切线空间
    vec3 up    = vec3(0.0, 1.0, 0.0);
    vec3 right = normalize(cross(up, N));
    up = normalize(cross(N, right));

    float sampleDelta = 0.025;
    float nrSamples = 0.0;
    for(float phi = 0.0; phi < 2.0 * PI; phi += sampleDelta)
    {
        for(float theta = 0.0; theta < 0.5 * PI; theta += sampleDelta)
        {
            // 球面到笛卡尔（在切线空间中）
            vec3 tangentSample = vec3(sin(theta) * cos(phi),  sin(theta) * sin(phi), cos(theta));
            // 切线空间到世界空间
            vec3 sampleVec = tangentSample.x * right + tangentSample.y * up + tangentSample.z * N;

            irradiance += texture(environmentMap, sampleVec).rgb * cos(theta) * sin(theta);
            nrSamples++;
        }
    }
    irradiance = PI * irradiance * (1.0 / float(nrSamples));

    FragColor = vec4(irradiance, 1.0);
}
```

创建一个 cube，将上述顶点着色器应用于该 cube，则 cube 的 6 个面就是辐照度环境立方体贴图的 6 个面

**·2.球谐函数 SH（球谐函数还没看懂，这里给出一些参考资料）**
参考：[球谐光照——球谐函数](https://zhuanlan.zhihu.com/p/153352797)
[Chapter 10. Real-Time Computation of Dynamic Irradiance Environment Maps（GPU Gems2）](https://developer.nvidia.com/gpugems/gpugems2/part-ii-shading-lighting-and-shadows/chapter-10-real-time-computation-dynamic)

**·3.球形高斯 SG**
参考：[SG Series Part 2: Spherical Gaussians 101](https://therealmjp.github.io/posts/sg-series-part-2-spherical-gaussians-101/)
球谐函数与球形高斯的原理与立方体贴图本质上差不多，是用不同的方法存储光照信息

### 环境光-镜面反射

镜面反射部分如下：
<img src="https://img-blog.csdnimg.cn/9f7ba1cd40434b30913e75fba208bea3.png" alt="在这里插入图片描述">
ks 实际上不是一个常数，它取决于入射光方向 ωi 和视线方向 v(ω0)（实际计算中因为 F 隐含了 ks 所以 ks 要删去）。同样因为对各个方向的 ωi 和 v 进行积分过于复杂所以无法实时计算。Epic 对此提出了一种解决方案，在做出一定妥协的情况下，为了实时计算的目的对镜面反射部分进行预卷积，这种方法被称为分解求和近似（Split Sum Approximation）
**分解求和近似**
该方面将镜面反射积分拆解成两个单独的积分。左边称为预过滤环境贴图（Pre-filtered environment map），右边则是 BRDF（环境光光 BRDF 描述镜面反射的部分），下面分别介绍。
<img src="https://img-blog.csdnimg.cn/5ddd17e5f704440fb429bd6c35087a21.png" alt="在这里插入图片描述">
等式的右边也可以表示成累加和的形式（也是我们之后用采样的方式求解积分的原理）
<img src="https://img-blog.csdnimg.cn/6741002488d94d64b184f164aeb4739c.png" alt="在这里插入图片描述">

1.预过滤环境贴图（Pre-filtered environment map）

<img src="https://img-blog.csdnimg.cn/aaf5892f40b444c9a612d25e27d1bb61.png" width="30%">
预过滤环境贴图跟辐照图贴图类似，是预先卷积计算过的环境贴图，但预过滤环境贴图考虑到了粗糙度（roughness）的影响。为了体现不同的粗糙度级别，环境贴图会使用更为分散的入射光向量（也叫采样向量）计算卷积以产生更模糊的镜面反射，我们将不同粗糙度的计算结果存入 mipmap 的不同级别中，如下图所示
<img src="https://img-blog.csdnimg.cn/eddd6d50152b431d97ed29cb38d3e1fc.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAd3NXaW5k,size_20,color_FFFFFF,t_70,g_se,x_16" alt="在这里插入图片描述">

**生成预过滤环境贴图**（重点：如何选择采样向量？）：
先引入几个概念（只是简单介绍，详细了解的话直接点击打开超链接）
[大数定律（Law of Large Numbers）](https://baike.baidu.com/item/%E5%A4%A7%E6%95%B0%E5%AE%9A%E5%BE%8B/410082?fr=aladdin) ：简单理解为一个概率为 P 的事件做 n 次实验，实验次数越多，事件发生的频率越接近于概率 P
[蒙特卡洛积分（Monte Carlo）](https://blog.csdn.net/hellocsz/article/details/94400402)：用一组满足分布律 p(x)的若干随机数对被积函数在积分区间内采样。蒙特卡洛积分建立在大数定律的基础上，式子如下
（h(x)是被积函数，p(x)是被积函数理论上的概率密度函数（PDF），N 是采样次数，设 f(x)=h(x)/p(x),则蒙特卡洛积分的推导和公式如下）
<img src="https://img-blog.csdnimg.cn/e7fca2c6b58b42e384f1b1d26da233c9.png">
最后一步根据大数定律知，n 越大，平均数越接近于期望

重要性采样（Importance Sampling）：蒙特卡洛积分选取采样点的一种方式。指选取的随机数（也叫采样点）集中在被积函数中对积分贡献较高的区域，而不是积分区间内均匀分布（直接采样）。这样选取可以减小方差，收敛速度快。下面给出两个参考链接：
[一文看懂蒙特卡洛采样方法](https://zhuanlan.zhihu.com/p/338103692)
[随机模拟-Monte Carlo 积分及采样](https://www.jianshu.com/p/3d30070932a8)
下图给出重要性采样的公式。
h(x)是被积函数，p(x)是被积函数实际上的概率密度函数，N 是采样次数，f(x)=h(x)/p(x)，q(x)是为了得出 p(x)而引入的自定义的概率密度函数,根据下图可知，p(x)/q(x)为重要性权值
<img src="https://img-blog.csdnimg.cn/cfa762104c934791b6e58bc86b0d2d72.png">
根据蒙特卡洛积分，最后一行的表达式就是我们要求的最终积分值
————
了解上面这些概念之后，我们再来看下面这幅图
<img src="https://img-blog.csdnimg.cn/3f6614734f7343728c84c49f3fe5a823.png" alt="在这里插入图片描述">
可以看到镜面反射的波瓣(lobe)尺寸随着粗糙度增加而变大。由于大多数光线都会集中在微平面法线向量 h 为中心的波瓣中，因此我们采样的光线（生成的样本向量）也应该满足这个规律（集中在波瓣中而不是像计算漫反射那样均匀取样本），这也是用到重要性采样的原因
<img src="https://img-blog.csdnimg.cn/aaf5892f40b444c9a612d25e27d1bb61.png" width="30%">
回看要积分的式子，下面给出 ue 中预过滤环境贴图的代码并给出解释（给定一个粗糙度 roughness 和镜面反射向量 R 求出该镜面反射方向上的辐亮度/颜色）
<img src="https://img-blog.csdnimg.cn/c76a8ccd47bd400e9c61c175d1c9b7b8.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAd3NXaW5k,size_20,color_FFFFFF,t_70,g_se,x_16" alt="在这里插入图片描述">
可以看到 epic 在此处令 N=V=R，这是因为在卷积预过滤环境贴图的时候我们并不知道视线方向 V。这意味着当我们从下图所示的角度观察时不会获得很好的镜面反射效果，但这通常被认为是合理的折衷方案，因此 fr 中 D 项（NDF)=1

<img src="https://img-blog.csdnimg.cn/f6d4000c40d54f569fd5b51466e49daf.png">

代码中 Hammersley 函数是生成随机数序列的一种方法，可以得到在[0,1]之间均匀分布的随机数序列（生成点在采样空间分布的均匀程度称作差异度 Discrepancy，Hammersley 属于低差异度序列），下面给出具体实现代码：(来自 learn opengl)

```c
float RadicalInverse_VdC(uint bits)
{
    bits = (bits << 16u) | (bits >> 16u);
    bits = ((bits & 0x55555555u) << 1u) | ((bits & 0xAAAAAAAAu) >> 1u);
    bits = ((bits & 0x33333333u) << 2u) | ((bits & 0xCCCCCCCCu) >> 2u);
    bits = ((bits & 0x0F0F0F0Fu) << 4u) | ((bits & 0xF0F0F0F0u) >> 4u);
    bits = ((bits & 0x00FF00FFu) << 8u) | ((bits & 0xFF00FF00u) >> 8u);
    return float(bits) * 2.3283064365386963e-10; // / 0x100000000
}
// ----------------------------------------------------------------------------
vec2 Hammersley(uint i, uint N)
{
    return vec2(float(i)/float(N), RadicalInverse_VdC(i));
}
```

重要性采样的函数如下，用于得到采样点且采样点分布受粗糙度影响（Phi:球面坐标系的 φ，H：将球面坐标转换为笛卡尔坐标系）：
<img src="https://img-blog.csdnimg.cn/496318f4b00240b28343b8788a754c34.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAd3NXaW5k,size_20,color_FFFFFF,t_70,g_se,x_16" alt="在这里插入图片描述">

最后 prefilteredColor 除以总样本权重，其中对最终结果影响较小的样本（NdotL 小）对最终结果的贡献就小。

**2.环境镜面反射 BRDF**
环境镜面反射与环境漫反射求解思路相同，也是对大量入射光进行采样来预计算光照信息
<img src="https://img-blog.csdnimg.cn/e408c65baae84c85b47a85ef18e282f4.png">
其中
<img src="https://img-blog.csdnimg.cn/1a0e040028a64b268265789660c42b1b.png">

这个卷积要求我们同时考虑到入射角（n·ω0），表面粗糙度（影响 fr 中 G 和 D 项）和菲涅尔系数 F0（与入射光矢量 ω0）,对 3 个变量卷积实在复杂，但是我们可以对方程做点变换，先试着将 F 项移出方程
<img src="https://img-blog.csdnimg.cn/fae95a1d023646e781aed0faee772cd0.png">
将分母移到 fr 下
<img src="https://img-blog.csdnimg.cn/7384f82b08324440afc1d02024d7cc2c.png">
F 项根据之前讲过的 Fresnel-Schlick 计算
<img src="https://img-blog.csdnimg.cn/f8f7e80bfbc84a019b9b4922399da8ad.png" alt="在这里插入图片描述">
为了式子看上去更简洁，我们用 α 表示（1-ω0·h）^5，之后再变形
<img src="https://img-blog.csdnimg.cn/d6575fb6b8cb4a64a6ad4870061af109.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAd3NXaW5k,size_19,color_FFFFFF,t_70,g_se,x_16" alt="在这里插入图片描述">
我们可以把最后一行式子拆成两项
<img src="https://img-blog.csdnimg.cn/09b80add91d045dca3d8174d39d1c47a.png" alt="在这里插入图片描述">
重新把 α 替换回来，把 fr/F 替换成一个新的 fr，这个新的 fr 已经不含菲涅尔项 F 了；同时我们把 F0 常数移到积分外面，最终式子如下
<img src="https://img-blog.csdnimg.cn/15ee749f99ae4f7999a87609f795bf46.png" alt="在这里插入图片描述">
这条函数形如 F0\*a+b,因此我们称左边积分为 F0_Scale（缩放系数）,右边积分为 F0_bias（偏移）。根据这条函数，我们就可以计算每一个像素（微元）的环境光镜面反射值，返回类型为 float2/vec2，我们以颜色的形式（scale 当作 r，bias 当作 g）存储在一张贴图上，用 roughness 和 NdotV 表示纹理索引/位置（这张贴图被称为 2D LUT/2D 查找纹理,也叫 BRDF 混合贴图（BRDF integration map））。示例图和代码实现如下（来自 learn opengl，UE 实现的方法是一样的）
<img src="https://img-blog.csdnimg.cn/545f6bd3629347f0b766bf7e6fdccc2a.png" alt="在这里插入图片描述">

```c
vec2 IntegrateBRDF(float NdotV, float roughness)
{
    vec3 V;
    V.x = sqrt(1.0 - NdotV*NdotV);
    V.y = 0.0;
    V.z = NdotV;

    float A = 0.0;
    float B = 0.0;

    vec3 N = vec3(0.0, 0.0, 1.0);

    const uint SAMPLE_COUNT = 1024u;
    for(uint i = 0u; i < SAMPLE_COUNT; ++i)
    {
        vec2 Xi = Hammersley(i, SAMPLE_COUNT);
        vec3 H  = ImportanceSampleGGX(Xi, N, roughness);
        vec3 L  = normalize(2.0 * dot(V, H) * H - V);//normalize(reflect(-V, H));

        float NdotL = max(L.z, 0.0);
        float NdotH = max(H.z, 0.0);
        float VdotH = max(dot(V, H), 0.0);

        if(NdotL > 0.0)
        {
            float G = GeometrySmith(N, V, L, roughness);
            float G_Vis = (G * VdotH) / (NdotH * NdotV);
            float Fc = pow(1.0 - VdotH, 5.0);

            A += (1.0 - Fc) * G_Vis;//A是F0_Scale
            B += Fc * G_Vis;//B是F0_Bias
        }
    }
    A /= float(SAMPLE_COUNT);
    B /= float(SAMPLE_COUNT);
    return vec2(A, B);
}


// ----------------------------------------------------------------------------
float GeometrySchlickGGX(float NdotV, float roughness)
{
    // G_ShclickGGX(N, V, k) = ( dot(N,V) ) / ( dot(N,V)*(1-k) + k )
    float a = roughness;
    float k = (a * a) / 2.0;

    float nom   = NdotV;
    float denom = (NdotV * (1.0 - k) + k) + 0.0001f;//防止分母为0

    return nom / denom;
}
// ----------------------------------------------------------------------------
float GeometrySmith(vec3 N, vec3 V, vec3 L, float roughness)
{
    float NdotV = max(dot(N, V), 0.0);
    float NdotL = max(dot(N, L), 0.0);
    float ggx2 = GeometrySchlickGGX(NdotV, roughness);
    float ggx1 = GeometrySchlickGGX(NdotL, roughness);

    return ggx1 * ggx2;
}

```

解释一下 G_Vis 的推导
G_Vis 根据公式其实指的就是去除 F 项后的 fr \* NdotL，
∴ G_Vis=G·D·NdotL / (4·NdotL · NdotV)=G·D/ (4· NdotV)，
加上之前提到假设 N=R=L 所以 D=1，
∴ G_Vis=G / (4 · NdotV)
根据重要性采样的公式，还需要乘上重要性权值（实际概率密度函数/自定义概率密度函数，这里自定义概率密度函数就是 Hammersley 法生成的[0,1]均匀分布序列，因此分母项为 1 可以省略，所以重要性权值就是实际概率密度函数，下文简称 PDF）然后求平均，这里的 PDF=4·VoH/NoH，与上式相乘结果与代码一致。(我是根据结果倒退的 PDF，PDF 等于这个值的时候就说的通了，这个 PDF 跟[ue 提到的 PDF](https://de45xmedrsdbp.cloudfront.net/Resources/files/2013SiggraphPresentationsNotes-26915738.pdf)正好互为倒数，我想了很久也没想通哪个是对的，还是说我推 G_Vis 的思路有问题，希望有大神能解答一下)

别忘了之前在介绍直接光 G 项的时候提到过，K 值与粗糙度 α 的关系，因此这里用 IBL 的方式计算环境光的时候，k=（α^2）/2
<img src="https://img-blog.csdnimg.cn/9f6a19b8b27744dfa04cf4ca7884c745.png" alt="在这里插入图片描述">
总结：在计算环境光 BRDF 项时，由于要对入射角（n·ω0，代码中的 NdotV），表面粗糙度 roughness（影响 fr 中 G 和 D 项）和菲涅尔系数 F0(与入射光矢量 ω0 有关)三个变量积分（补充：尽管 roughness 与被积微元 dω0 没关系，但是每一个像素可能都有不同的粗糙度，因此在积分中也要考虑不同 roughness 的影响）过于复杂，因此采用预计算的方式，得出不同入射角、不同粗糙度下 F0 的值，将这些信息存储在一张贴图（称为 LUT 查找纹理）中。
在使用时，我们用 NdotV 和 roughness 对这张 LUT 采样，最终环境光 BRDF 的值为 F0\*LUT.r+LUR.g
————————
得到预过滤环境贴图与环境光 BRDF 项后，将两者相乘得到环境光-镜面反射值（这里的 SpecularColor 就是指 F0），最终基于 IBL 的环境光-镜面反射的代码如下<img src="https://img-blog.csdnimg.cn/fcd075faa2ca4d5382b624722adee6a9.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAd3NXaW5k,size_20,color_FFFFFF,t_70,g_se,x_16" alt="在这里插入图片描述">

将环境光漫反射与镜面反射相加得到完整的环境光，这里可以乘上一个系数 AO 来模拟物体之间挨的很近时环境光很弱产生的阴影（AO 具体怎么计算本文不做过多叙述，可以参考网上其他资料，这里给出一个我觉觉得讲的挺好的：[游戏中的全局光照(三) 环境光遮蔽/AO](https://zhuanlan.zhihu.com/p/194198670)）

最终将直接光与环境光相加得到完整的光照模型

## 考虑次表面散射的 BSSRDF（待补充）

参考资料：[游戏中的次表面散射/Subsurface Scattering](https://zhuanlan.zhihu.com/p/337384739)

**2.手写一个基于能量守恒的光照**
代码如下：（间接光部分抄 unity 源码，但是间接光高光反射似乎不太正确）

```c
//unity2019.4.28f1c1，不同版本可能源代码位置不同
Shader "Custom/CustomPBR"
{

    Properties
    {
        _BaseColor ("Basecolor", Color) = (1, 1, 1, 1)
        _Albedo ("Albedo", 2D) = "white" { }
        [NoScaleOffset][Normal]_Normal ("Normal", 2D) = "bump" { }
        //[NoScaleOffset]_AO ("AO", range(0, 1)) = 0
        [NoScaleOffset]_MaskMap ("Mask", 2D) = "white" { }
        //[NoScaleOffset]_Roughness ("Roughness", 2D) = "white"{}
        _Roughness ("perceptualRoughness", range(0, 1)) = 1 //被人体直观感知到的线性变化的粗糙度
        _Bumpscale ("Bumpscale", range(0, 1)) = 1
        _Metallic ("Metallic", range(0, 1)) = 1
    }
    SubShader
    {

        Pass
        {

            Tags { "LightMode" = "UniversalForward" }
            HLSLPROGRAM

            #pragma vertex vert
            #pragma fragment frag

            #include "Packages/com.unity.render-pipelines.universal/ShaderLibrary/Core.hlsl"
            #include "Packages/com.unity.render-pipelines.universal/ShaderLibrary/Lighting.hlsl"

            CBUFFER_START(UnityPerMaterial)
                float4 _Albedo_ST;
                float4 _Normal_ST;
                float _Bumpscale;
                float _Metallic;
                float4 _BaseColor;
                float _Roughness;
                //float _AO;
            CBUFFER_END
            sampler2D _Albedo;
            sampler2D _Normal;
            Texture2D _MaskMap;
            SAMPLER(sampler_MaskMap);

            //用于直接光计算

            //法线分布函数
            float DistributionGGX(float NdotH, float roughness)
            {
                float a = roughness * roughness;
                float a2 = a * a;//分子
                float denom = NdotH * NdotH * (a2 - 1) + 1;//分母
                denom = denom * denom * PI;
                return a2 / denom;
            }

            //菲涅尔方程
            float3 FresnelSchlick(float3 F0, float VdotH)
            {
                //return F0+(1-F0)*pow(1-VdotH,5);
                return F0 + (1 - F0) * exp2((-5.55473 * VdotH - 6.98316) * VdotH);//ue4 in 2013siggraph，unity进一步用vdoth代替vdot

            }

            float GeometrySchlickGGX(float NdotV, float roughness)
            {
                float k = (roughness + 1) * (roughness + 1) / 8;//h_direct
                float nom = NdotV;
                float denom = NdotV * (1.0 - k) + k;
                return nom / denom;
            }

            //阴影遮罩函数
            float GeometrySmith(float NdotV, float NdotL, float roughness)
            {
                float ggx1 = GeometrySchlickGGX(NdotV, roughness);
                float ggx2 = GeometrySchlickGGX(NdotL, roughness);
                return ggx1 * ggx2;
            }

            //用于间接光

            float3 FresnelSchlickRoughness(float NdotV, float3 F0, float roughness)
            {
                //return F0 + saturate(1 - roughness - F0) * pow(clamp(1.0 - NdotV, 0.0, 1.0), 5.0);
                return F0 + saturate(1 - roughness - F0) * exp2((-5.55473 * NdotV - 6.98316) * NdotV);//拟合

            }

            //从unity_SpecCube0采样
            float3 MyGlossyEnvironmentReflection(half3 normalWS, float3 viewWS, half perceptualRoughness, half AO)//line 589 in Lighting.hlsl

            {
                float3 reflectVector = reflect(-viewWS, normalWS);
                return GlossyEnvironmentReflection(reflectVector, perceptualRoughness, AO);
            }

            //得到反射率
            half3 MyReflectivitySpecular(half3 specular)//line 270 in Lighting.hlsl

            {
                #if defined(SHADER_API_GLES)
                    return specular.r;

                #else
                    return max(max(specular.r, specular.g), specular.b);
                #endif
            }

            half3 MyEnvironmentBRDFSpecular(float roughness2, float smoothness, half3 F0, float NdotV)//line 371 in Lighting.glsl

            {
                half fresnelTerm = Pow4(1.0 - NdotV);
                float surfaceReduction = 1.0 / (roughness2 * roughness2 + 1.0);

                float reflectivity = MyReflectivitySpecular(F0);
                float grazingTerm = saturate(smoothness + reflectivity);

                return surfaceReduction * lerp(F0, grazingTerm, fresnelTerm);
            }


            struct a2v
            {
                float4 positionOS : POSITION;
                float4 normalOS : NORMAL;
                float4 tangentOS : TANGENT;
                float4 uv : TEXCOORD0;
            };
            struct v2f
            {
                float4 positionCS : SV_POSITION;
                float4 uv : TEXCOORD0;
                float4 normalWS : NORMAL;
                float4 tangentWS : TANGENT;
                float4 biotangentWS : TEXCOORD1;
            };

            v2f vert(a2v v)
            {
                v2f o;
                o.positionCS = TransformObjectToHClip(v.positionOS.xyz);
                float3 positionWS = TransformObjectToWorld(v.positionOS.xyz);
                o.normalWS.xyz = normalize(TransformObjectToWorldNormal(v.normalOS.xyz));
                o.tangentWS.xyz = normalize(TransformObjectToWorldDir(v.tangentOS.xyz));
                o.biotangentWS.xyz = normalize(cross(o.normalWS.xyz, o.tangentWS.xyz) * v.tangentOS.w);
                o.normalWS.w = positionWS.x;
                o.tangentWS.w = positionWS.y;
                o.biotangentWS.w = positionWS.z;
                o.uv.xy = TRANSFORM_TEX(v.uv, _Albedo);
                o.uv.zw = v.uv;
                return o;
            }

            half4 frag(v2f i) : SV_Target
            {
                //提取mask贴图中的金属度，AO和粗糙度
                float4 Mask = SAMPLE_TEXTURE2D(_MaskMap, sampler_MaskMap, i.uv.zw);
                float Metallic = Mask.r;//Mask.r/_Metallic
                float AO = Mask.g;
                float Roughness = _Roughness;//这里的Roughness等价于unity源码中的perceptualSmoothness
                //这里的Roughness2等价于unity源码中的roughness,Roughness平方主要是考虑到人眼对粗糙度的感知是非线性的
                float Roughness2 = Roughness * Roughness;
                float smoothness = 1 - _Roughness;///Mask.a

                Light light = GetMainLight();
                half3 Clight = light.color;
                float3 L = normalize(light.direction);
                float3 positionWS = float3(i.normalWS.w, i.tangentWS.w, i.biotangentWS.w);
                float3 V = SafeNormalize(_WorldSpaceCameraPos.xyz - positionWS);
                float3 H = SafeNormalize(L + V);
                float3 Albedo = tex2D(_Albedo, i.uv.xy).xyz * _BaseColor.xyz;
                float3 F0 = lerp(float3(0.04, 0.04, 0.04), Albedo, Metallic);
                //计算法线（把法线从贴图的切线空间转换到世界空间下）
                float3x3 TtoW = {
                    i.tangentWS.xyz, i.biotangentWS.xyz, i.normalWS.xyz
                };

                TtoW = transpose(TtoW);
                half3 NormalTS = UnpackNormalScale(tex2D(_Normal, i.uv.zw), _Bumpscale);
                // half3 NormalTS = UnpackNormal(tex2D( _Normal, i.uv.zw));
                // NormalTS.xy*=_Bumpscale;
                NormalTS.z = sqrt(1 - saturate(dot(NormalTS.xy, NormalTS.xy)));
                float3 N = normalize(mul(TtoW, NormalTS));


                //预先计算必要的点积
                float NdotH = max(dot(N, H), 0.000001);
                float VdotH = max(dot(V, H), 0.000001);
                float NdotV = max(dot(N, V), 0.000001);
                float NdotL = max(dot(N, L), 0.000001);

                //直接光漫反射
                float3 diffuse = Albedo;
                float3 Direct_Diffuse = diffuse * Clight * NdotL;//除于Π和最后的球面积分Π正好消掉

                //直接光高光反射
                float D = DistributionGGX(NdotH, Roughness);
                float3 F = FresnelSchlick(F0, VdotH);
                float G = GeometrySmith(NdotV, NdotL, Roughness);
                float3 specular = 0.25 * D * F * G / (NdotV * NdotL);
                float3 Direct_Specular = specular * Clight * NdotL;

                //直接光
                float3 ks = F;
                float3 kd = (1 - ks) * (1 - Metallic);
                //高光反射不用乘ks，因为本身就已经带了菲涅尔项F;
                //也不需要乘上Π，因为只有L+V=H的方向光线才会进入眼睛，不像漫反射会接收到来自四面八方的光，所以不需要对球面积分
                float3 DirectColor = kd * Direct_Diffuse + Direct_Specular;

                //间接光漫反射
                float3 SHColor = SampleSH(N);
                float3 Indir_Diffuse = SHColor * Albedo;

                //间接光高光反射
                float3 reflectDir = reflect(-V, N);
                float3 IndirSpeEnvColor = MyGlossyEnvironmentReflection(N, V, Roughness2, AO);
                float3 IndirSpeFactor = MyEnvironmentBRDFSpecular(Roughness2, smoothness, F0, NdotV);
                float3 Indir_Specular = IndirSpeFactor * IndirSpeEnvColor;

                //间接光
                float3 kS = FresnelSchlickRoughness(NdotV, F0, Roughness);
                float3 kD = (1.0 - kS) * (1 - Metallic);
                float3 IndirectColor = kD * Indir_Diffuse * AO + Indir_Specular;

                //合并直接光和间接光
                float3 output = DirectColor + IndirectColor;

                return float4(output, 1.0f);
            }
            ENDHLSL
        }
    }
}
```

效果图（在线性空间下渲染，左边是自己写的 shader，右边是官方的 lit shader，虽然在纯色材质下 差距有点大，不过差强人意）：
<img src="https://img-blog.csdnimg.cn/74055c4bdf9d4da4afcb07310e852ddb.png" alt="在这里插入图片描述">
（注:上图测试的时候使用 linear 空间，albedo 需要勾选 srgb)
