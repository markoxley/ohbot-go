package ohbot

const (
	xmlDefault = `<SettingList>
    <Setting Name="DefaultEyeShape" Value="eyeball" /> 
    <Setting Name="DefaultSpeechSynth" Value="" /> 
    <Setting Name="DefaultVoice" Value="" /> 
    <Setting Name="DefaultLang" Value="en-gb" /> 
    <Setting Name="SpeechDBFile" Value="ohbotData/OhbotSpeech.csv" /> 
    <Setting Name="EyeShapeList" Value="ohbotData/ohbot.obe" /> 
    <Setting Name="MotorDefFile" Value="ohbotData/MotorDefinitionsv21.omd" />
</SettingList>`

	motorDef = `<Motors>
  <Motor Name="HeadTurn" Min="0" Max="1000" Motor="1" Speed="40" Reverse="False" Acceleration="60" RestPosition="5" Avoid="" />
  <Motor Name="HeadNod" Min="140" Max="700" Motor="0" Speed="0" Reverse="True" Acceleration="60" RestPosition="5" Avoid="" />
  <Motor Name="EyeTurn" Min="380" Max="780" Motor="2" Speed="0" Reverse="False" Acceleration="0" RestPosition="5" Avoid="" />
  <Motor Name="EyeTilt" Min="520" Max="920" Motor="6" Speed="0" Reverse="False" Acceleration="30" RestPosition="5" Avoid="" />
  <Motor Name="TopLip" Min="0" Max="550" Motor="4" Speed="0" Reverse="True" Acceleration="0" RestPosition="5" Avoid="BottomLip" />
  <Motor Name="BottomLip" Min="0" Max="550" Motor="5" Speed="0" Reverse="True" Acceleration="0" RestPosition="5" Avoid="TopLip" />
  <Motor Name="LidBlink" Min="35" Max="305" Motor="3" Speed="0" Reverse="False" Acceleration="0" RestPosition="10" Avoid="" />
  <Motor Name="MouthOpen" Min="80" Max="460" Motor="7" Speed="0" Reverse="False" Acceleration="0" RestPosition="10" Avoid="" />
</Motors>`

	eyeDef = `<EyeShapeSet xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <EyeShapeList>
    <EyeShape>
      <Name>Angry</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>5</PupilRangeX>
      <PupilRangeY>5</PupilRangeY>
      <Matrix/>
      <Hex>2070F8FCFE7C3800000070F8FCFE3C000000000078FCFC38000000000000F8FC0000000000000000F800000000000010381000000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
    <EyeShape>
      <Name>BoxLeft</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>6</PupilRangeX>
      <PupilRangeY>7</PupilRangeY>
      <Matrix/>
      <Hex>00001F1F1F1F1F00000000000004000000000000000E0A0E00000000001F1115111F00000000000E0A0E000000000000000C0C000000</Hex>
      <AutoMirror>false</AutoMirror>
    </EyeShape>
    <EyeShape>
      <Name>BoxRight</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>6</PupilRangeX>
      <PupilRangeY>7</PupilRangeY>
      <Matrix/>
      <Hex>007C7C7C7C7C7C0000000000202000000000000038282838000000007C445454447C0000000038282838000000000000003030000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
    <EyeShape>
      <Name>Crying</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>5</PupilRangeX>
      <PupilRangeY>5</PupilRangeY>
      <Matrix/>
      <Hex>003C7EFEFEFC002000003C7EFEFE7C800020387CFEFEFE7C388800387CFEFEFE7C380088003C7EFEFEFC082800000010381000000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
    <EyeShape>
      <Name>Eyeball</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>5</PupilRangeX>
      <PupilRangeY>5</PupilRangeY>
      <Matrix/>
      <Hex>387CFEFEFE7C380000007CFEFEFE7C00000000007CFE7C00000000000000FE7C00000000000000827C00000000000010381000000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
    <EyeShape>
      <Name>Glasses</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>6</PupilRangeX>
      <PupilRangeY>5</PupilRangeY>
      <Matrix/>
      <Hex>FEFFFEFEFE00000000FEFFFEFEFE00000000FEFFFEFEFE00000000FEFFFEFEFE00000000FEFFFEFEFE00000000001038100000000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
    <EyeShape>
      <Name>Heart</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>5</PupilRangeX>
      <PupilRangeY>6</PupilRangeY>
      <Matrix/>
      <Hex>6CFEFEFE7C38100000007CFEFE7C3810000000007CFE7C38100000000000FE7C38000000000000007C00000000000010381000000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
    <EyeShape>
      <Name>Large</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>5</PupilRangeX>
      <PupilRangeY>5</PupilRangeY>
      <Matrix/>
      <Hex>387CFEFEFE7C380000007CFEFEFE7C00000000007CFE7C00000000000000FE7C00000000000000827C000000000010387C3810000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
    <EyeShape>
      <Name>Sad</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>5</PupilRangeX>
      <PupilRangeY>5</PupilRangeY>
      <Matrix/>
      <Hex>081C3E7EFE7C380000001C3E7EFE7C00000000003C7E7E3C0000000000003E7E00000000000000003E00000000000010381000000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
    <EyeShape>
      <Name>SmallBall</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>5</PupilRangeX>
      <PupilRangeY>5</PupilRangeY>
      <Matrix/>
      <Hex>00387C7C7C3800000000107C7C7C100000000000387C3800000000000000380000000000000000000000000000000000100000000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
    <EyeShape>
      <Name>Square</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>5</PupilRangeX>
      <PupilRangeY>5</PupilRangeY>
      <Matrix/>
      <Hex>7CFEFEFEFEFE7C0000007CFEFEFE7C00000000007CFE7C00000000000000FE7C00000000000000827C00000000000010381000000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
    <EyeShape>
      <Name>SunGlasses</Name>
      <FileName>Custom.obe</FileName>
      <PupilRangeX>0</PupilRangeX>
      <PupilRangeY>0</PupilRangeY>
      <Matrix/>
      <Hex>FE838282FE00000000FEC78282FE00000000FEFF8282FE00000000FEFFC682FE00000000FEFFFE82FE00000000000000000000000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
    <EyeShape>
      <Name>VerySad</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>6</PupilRangeX>
      <PupilRangeY>6</PupilRangeY>
      <Matrix/>
      <Hex>1C3E7EFEFEFE7C0000003C7EFEFE7C00000000007CFE7C0000000000007EFC00000000000000827C000000000000183C3C3C18000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
  <EyeShape>
      <Name>Full</Name>
      <FileName>Ohbot.obe</FileName>
      <PupilRangeX>5</PupilRangeX>
      <PupilRangeY>5</PupilRangeY>
      <Matrix/>
      <Hex>ffffffffffffffffffffffffffffffff0000ffffffff0000000000ffff00000000000000000000000000000000000000000000000000</Hex>
      <AutoMirror>true</AutoMirror>
    </EyeShape>
    </EyeShapeList>
  <Version>1</Version>
  <FilePathName/>
</EyeShapeSet>`

	speechDef = `Set,Variable,Phrase
1,1,what is your command?
1,2,how can I help?
1,3,what's up?
1,4,ready for input
2,1,Let me think
2,2,Just a second
2,3,give me a moment
2,4,let me get you an answer
3,1,Hello my name is Ohbot`
)
