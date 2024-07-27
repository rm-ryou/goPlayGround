package model

type HittableList struct {
	Objects []Sphere
}

func (hl *HittableList) AddObject(sphere Sphere) {
	hl.Objects = append(hl.Objects, sphere)
}

func (hl *HittableList) Hit(ray Ray, rayTMax, rayTMin float64, rec *HitRecord) bool {
	tmpRec := new(HitRecord)
	isHit := false
	closestToFar := rayTMax

	for _, object := range hl.Objects {
		if object.Hit(ray, closestToFar, rayTMin, tmpRec) {
			isHit = true
			closestToFar = tmpRec.T
			*rec = *tmpRec
		}
	}
	return isHit
}
