package file

type Info struct {

}

func GetInfo(path string) (Info, error) {
	_, err := os.Stat(path)
  if err != nil {
    if os.IsNotExist(err) {
      return false, nil
    } else {
      return false, err
    }
  }
  return true, nil
}


// Exists checks whether the given file exists or not.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
  if err != nil {
    if os.IsNotExist(err) {
      return false, nil
    } else {
      return false, err
    }
  }
  return true, nil
}

func IsFile(path string) (bool, error) {
  fi, err :=  os.Stat(path)
  if err != nil {
    return false, err
  }
  return !fi.Mode().IsDir(), err
}

func IsDir(path string) (bool, error) {
  fi, err :=  os.Stat(path)
  if err != nil {
    return false, err
  }
  return fi.Mode().IsDir(), err
}

func IsRegular(path string) (bool, error) {
  fi, err :=  os.Stat(path)
  if err != nil {
    return false, err
  }
  return fi.Mode().IsRegular(), err
}

func isReadable(path string) (bool, error) {
  fi, err :=  os.Stat(path)
  if err != nil {
    return false, err
  }
  return fi.Mode() & 0400 != 0 , err
}

func isWritable(path string) (bool, error) {
  fi, err :=  os.Stat(path)
  if err != nil {
    return false, err
  }
  return fi.Mode() & 0200 != 0 , err
}

func isExecutable(path string) (bool, error) {
  fi, err :=  os.Stat(path)
  if err != nil {
    return false, err
  }
  return fi.Mode() & 0100 != 0 , err
}
