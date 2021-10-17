class TimeTable {
  constructor(parameters) {
    
  }
}

    // outputTimeTable({ bandCount, rehearsalTime, faceToFaceMeeting, performanceTime, ventilation, startTime }) {
    //   this.timeTable = "";
    //   // バンドの配列作成 ex) bandCountが5なら ["バンド1", "バンド2", "バンド3", "バンド4", "バンド5"]
    //   const bands = [...Array(Number(bandCount)).keys()].map(i => `バンド${++i}`);

    //   this.time = this.$moment(`2021-01-01T${startTime}:00`);
    //   this.rehearsal(bands, rehearsalTime, ventilation);
    //   this.performance_preparation(faceToFaceMeeting);
    //   this.performance(bands, performanceTime, ventilation);
    //   this.$modal.show("modal-content");
    //   return;
    // },
    // rehearsal(bands, rehearsalTime, ventilation) {
    //   // 最初にだけ音出しバンドが存在する
    //   this.timeTable += `${this.time.format('HH:mm')}〜${this.time.add(rehearsalTime, 'm').format('HH:mm')} 音出しバンド\n`;

    //   // 換気のために別でカウントする
    //   let isVentilationCount = 2;
    //   for (var i = bands.length - 1;i > -1;i--) {
    //     // タイムテーブルの肝の部分 ex) 15:00〜15:15 バンド1
    //     this.timeTable += `${this.time.format('HH:mm')} 〜 ${this.time.add(rehearsalTime, 'm').format('HH:mm')} ${bands[i]}\n`;

    //     // 転換分で5分追加。顔合わせはすぐやるため最後のバンドの時は追加しない。
    //     if (ventilation && isVentilationCount % 3 == 0 && i !== 0) {
    //       this.timeTable += `${this.time.format('HH:mm')} 〜 ${this.time.add(5, 'm').format('HH:mm')} <換気>\n`;
    //     } else if (i !== 0) {
    //     // リハの転換
    //     //  this.time.add(5, 'm'); 
    //     }
    //     isVentilationCount++
    //   }
    //   return;
    // },
    // performance_preparation(faceToFaceMeeting) {
    //   this.timeTable += `${this.time.add(10, 'm').format('HH:mm')} ＼＼＼\\顔合わせ//／／／\n`
      
    //   // 顔合わせ後開ける時間
    //   this.timeTable += `START  [[[   ${this.time.add(faceToFaceMeeting, 'm').format('HH:mm')}   ]]]\n`
    //   return;
    // },
    // performance(bands, performanceTime, ventilation) {
    //   // 換気のために別でカウントする
    //   let isVentilationCount = 1;
    //   for (var i = 0;i < bands.length;i++) {
    //     // タイムテーブルの肝の部分 ex) 15:00〜15:15 バンド1
    //     this.timeTable += `${this.time.format('HH:mm')} 〜 ${this.time.add(performanceTime, 'm').format('HH:mm')} ${bands[i]}\n`;

    //     // 転換分で10分追加
    //     if (ventilation && isVentilationCount % 3 == 0 && i !== bands.length -1) {
    //       this.timeTable += `${this.time.format('HH:mm')} 〜 ${this.time.add(5, 'm').format('HH:mm')} <換気>\n`;
    //     } else {
    //       this.time.add(10, 'm')
    //     }
    //     isVentilationCount++
    //   }
    //   return;
    // }