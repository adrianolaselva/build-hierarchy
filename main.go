package main

type WorkerInPut struct {
	Id       int
	Name     string
	ParentId int
}

type WorkerOutPut struct {
	Id       int
	Name     string
	Children []WorkerOutPut
}

func BuildHierarchy(id int, workers []*WorkerInPut) WorkerOutPut {
	var out = WorkerOutPut{}
	for _, w := range workers {
		if id != w.Id {
			continue
		}

		out.Id = w.Id
		out.Name = w.Name
		BuildChild(workers, &out)
	}

	return out
}

func BuildChild(in []*WorkerInPut, out *WorkerOutPut) {
	for _, w := range in {
		if w.ParentId != out.Id {
			continue
		}

		out.Children = append(out.Children, BuildHierarchy(w.Id, in))
	}
}
