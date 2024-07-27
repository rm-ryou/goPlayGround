package model

type HittableList struct {
	Objects []Sphere
}

func (hl *HittableList) AddObject(sphere Sphere) {
	hl.Objects = append(hl.Objects, sphere)
}

func (hl *HittableList) Hit(ray Ray, rayT Interval, rec *HitRecord) bool {
	tmpRec := new(HitRecord)
	isHit := false
	closestToFar := rayT.Max

	for _, object := range hl.Objects {
		if object.Hit(ray, Interval{closestToFar, rayT.Min}, tmpRec) {
			isHit = true
			closestToFar = tmpRec.T
			*rec = *tmpRec
		}
	}
	return isHit
}
