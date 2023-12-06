package handler

// func (s *Server) registerJob1() error {

// 	_, err := s.cronjob.AddFunc("@every 0h0m10s", func() {
// 		log.Println("job 1 running....")
// 	})
// 	return err

// }

func (s *Server) registerJob2() error {

	_, err := s.cronjob.AddFunc("50 * * * *", func() {
		MainJob(s)
	})
	return err

}

// func (s *Server) mainJob2() {

// 	contentIdLength := 7122
// 	nJob := 10
// 	nSecond := 60

// 	res := util.CalcIndexJob(contentIdLength, nJob)

// 	for index, item := range res {
// 		now := time.Now().Add(time.Minute)
// 		t := now.Add(time.Second * time.Duration(nSecond*index))

// 		spec := fmt.Sprintf("%d * * * *", t.Minute())

// 		log.Println(spec)

// 		cronId, err := s.cronjob.AddFunc(spec, func(index int, item []int) func() {
// 			return func() {
// 				log.Printf("Executing job for index %d at time %v", index, time.Now())
// 				// job.CallApiContentAndInsertDB(item[0], item[1])
// 			}
// 		}(index, item))
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		log.Printf("Scheduled job %d with ID %d at %v:[%d %d]", index, cronId, t, item[0], item[1])
// 		time.Sleep(time.Second)
// 	}
// 	s.cronjob.Start()

// }
